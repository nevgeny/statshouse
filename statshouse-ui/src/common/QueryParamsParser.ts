// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import { dequal } from 'dequal/lite';
import { KeysTo, TIME_RANGE_KEYS_TO } from './TimeRange';
import { filterInSep, filterNotInSep, queryValueBackendVersion1, queryValueBackendVersion2 } from '../view/api';
import produce from 'immer';
import { deepClone } from '../view/utils';

const maxPrefixArray = 100;
const removeValueChar = String.fromCharCode(7);

export type ConfigParam<T = any, T2 = T> = {
  always?: boolean;
  /**
   * simple prefix
   */
  prefix?: string;
  /**
   * callback prefix if isArray=true
   * @param index - index item
   */
  prefixArray?: (index: number) => string;
  /**
   * param as array
   */
  isArray?: boolean;
  /**
   * param as object
   */
  fromEntries?: boolean;
  /**
   * list param in object
   */
  params?: Record<string, ConfigParam>;
  /**
   * default value
   */
  default?: T2;
  /**
   * name get param in url
   */
  urlKey?: string;
  /**
   * required param of isArray item
   * if not then end array
   */
  required?: boolean;
  /**
   * struct param of isArray item
   * if change then rewrite all
   */
  struct?: boolean;
  /**
   * encode get param
   * @param value
   */
  encode?: (value: T) => string | undefined;
  /**
   * decode get param
   * @param value
   */
  decode?: (value: string) => T2 | undefined;
};
export type ConfigParams = Record<string, ConfigParam>;

function getDecode<T, T2 = T>(values: string[], config: ConfigParam<T, T2>): T2 | (T2 | undefined)[] | undefined {
  const { decode = (s) => s as T2 } = config;
  if (values.length === 0) {
    return config.default;
  }
  if (config.isArray) {
    return values.map(decode);
  }
  if (config.fromEntries) {
    return values.map(decode);
  }
  if (config.required && values[0] === removeValueChar) {
    return values[0] as T2;
  }
  return decode(values[0]) ?? config.default;
}

function getEncode<T, T2 = T>(value: T | T[], config: ConfigParam<T, T2>): (string | undefined)[] | undefined {
  const { encode = (s: T) => s as string, always } = config;
  if (typeof value === 'undefined') {
    return [undefined];
  }
  if (!always && !config.fromEntries && dequal(value, config.default)) {
    return undefined;
  }
  if (!always && config.fromEntries && dequal((value as [string, unknown])[1], config.default)) {
    return undefined;
  }
  if (config.isArray && !config.fromEntries && isArray(value)) {
    return value.map(encode);
  }
  return [encode(value as T)];
}

function isArray(item: unknown): item is unknown[] {
  return typeof item === 'object' && item !== null && Array.isArray(item);
}

function isObject(item: unknown): item is Record<string, unknown> {
  return typeof item === 'object' && item !== null && !Array.isArray(item);
}

export function mergeLeft<T>(targetMerge: T, valueMerge: T): T {
  if (targetMerge === valueMerge) {
    return targetMerge;
  }
  if (isArray(targetMerge) && isArray(valueMerge)) {
    if (targetMerge.length === valueMerge.length) {
      return produce(targetMerge, (s) => {
        for (let i = 0, max = s.length; i < max; i++) {
          const v = mergeLeft(s[i], valueMerge[i]);
          if (s[i] !== v) {
            s[i] = v;
          }
        }
      });
    }
    return valueMerge;
  }
  if (isObject(targetMerge) && isObject(valueMerge)) {
    const tKey = Object.keys(targetMerge);
    const vKey = new Set(Object.keys(valueMerge));
    return produce(targetMerge, (s) => {
      tKey.forEach((key) => {
        const v = mergeLeft(s[key], valueMerge[key]);
        if (!valueMerge.hasOwnProperty(key)) {
          delete s[key];
        } else if (s[key] !== v) {
          Object.assign(s, { [key]: v });
        }
        vKey.delete(key);
      });
      [...vKey].forEach((key) => {
        Object.assign(s, { [key]: valueMerge[key] });
      });
    });
  }
  return valueMerge;
}

/**
 * decode/encode number param
 */
export const NumberParam: ConfigParam<number | undefined> = {
  encode: (s) => s?.toString(),
  decode: (s) => {
    if (s && !isNaN(+s)) {
      return +s;
    }
    return undefined;
  },
};

/**
 * decode/encode JSON object param
 */
export const ObjectsParam: (separator: string) => ConfigParam<[string, unknown]> = (separator) => ({
  fromEntries: true,
  encode: ([key, value]) => `${key}${separator}${typeof value === 'string' ? value : JSON.stringify(value)}`,
  decode: (s) => {
    const [key, value] = s.split(separator, 2);
    if (typeof value === 'undefined') {
      return undefined;
    }
    try {
      return [key, JSON.parse(value)];
    } catch (_) {
      return [key, value];
    }
  },
});

/**
 * decode/encode boolean param
 */
export const BooleanParam: ConfigParam<boolean | undefined> = {
  encode: (s) => (s ? '1' : '0'),
  decode: (s) => s === '1',
};

/**
 * decode/encode time to param
 */
export const TimeToParam: ConfigParam<number | KeysTo | undefined> = {
  encode: (s) => s?.toString(),
  decode: (s) => {
    if (s && !isNaN(+s)) {
      return +s;
    }
    if (s && Object.values(TIME_RANGE_KEYS_TO).includes(s as KeysTo)) {
      return s as KeysTo;
    }
    return undefined;
  },
};

/**
 * decode/encode tag sync param
 */
export const TagSyncParam: ConfigParam<(number | undefined)[]> = {
  isArray: true,
  decode: (s) => [
    ...s.split('-').reduce((res, t) => {
      const [plot, tagKey] = t.split('.').map((r) => parseInt(r));
      res[plot] = tagKey;
      return res;
    }, [] as (number | undefined)[]),
  ],
  encode: (v) =>
    v
      .map((key, index) => (Number.isInteger(key) ? `${index}.${key}` : undefined))
      .filter((s) => s)
      .join('-'),
};

/**
 * decode/encode filter param
 */
export const FilterParams: (notIn?: boolean) => ConfigParam<[string, string]> = (notIn) => ({
  isArray: true,
  fromEntries: true,
  decode: (s) => {
    const pos = s.indexOf(notIn ? filterNotInSep : filterInSep);
    const pos2 = s.indexOf(notIn ? filterInSep : filterNotInSep);
    if (pos === -1 || (pos > pos2 && pos2 > -1)) {
      return undefined;
    }
    const indexTag = s.substring(0, pos).replace('skey', '_s').replace('key', '');
    const tagID = '_s' === indexTag ? 'skey' : `key${indexTag}`;
    const tagValue = s.substring(pos + 1);
    return [tagID, tagValue];
  },
  encode: (v) => {
    const [tagID, tagValue] = v;
    const indexTag = tagID.replace('key', '').replace('s', '_s');
    return `${indexTag}${notIn ? filterNotInSep : filterInSep}${tagValue}`;
  },
});

/**
 * decode/encode v2 param
 */
export const UseV2Param: ConfigParam<boolean> = {
  encode: (s) => (s ? queryValueBackendVersion2 : queryValueBackendVersion1),
  decode: (s) => s === queryValueBackendVersion2,
};

function valueToArray<T extends Record<string, unknown>>(
  configParams: ConfigParams,
  value: T,
  defaultParams?: T,
  urlSearchParams: URLSearchParams = new URLSearchParams(window.location.search)
): [string, string | undefined][] {
  return Object.entries(configParams).flatMap(([key, config]) => {
    const { urlKey, prefix = '', prefixArray, params, fromEntries } = config;
    const nameParam = prefix + (urlKey ?? key);
    if (value) {
      if (config.isArray && prefixArray && params) {
        if (!isArray(value[key])) {
          return [[nameParam, undefined]];
        }
        const _default = (defaultParams?.[key] as unknown[] | undefined) ?? config.default?.[key];
        if (dequal(value[key], _default)) {
          return [[nameParam, undefined]];
        }
        return [...(value[key] as Record<string, unknown>[]), undefined].flatMap((v, index) => {
          let prefixA = prefixArray(index);
          const itemConfig = Object.fromEntries(
            Object.entries(params).map(([pKey, pConfig]) => [
              pKey,
              { ...pConfig, prefix: prefixA + (pConfig.prefix ?? '') },
            ])
          );
          const _default = config.struct
            ? config.default?.[index]
            : (defaultParams?.[key] as unknown[] | undefined)?.[index] ?? config.default?.[index];
          return valueToArray(itemConfig, v, _default, urlSearchParams);
        });
      } else if (fromEntries) {
        if (!isObject(value[key])) {
          return [[nameParam, undefined]];
        }
        if (dequal(value[key], defaultParams?.[key])) {
          return [[nameParam, undefined]];
        }
        if (defaultParams?.[key] && !Object.keys(value[key] as Record<string, unknown>).length) {
          return [[nameParam, removeValueChar]];
        }

        return Object.entries(value[key] as Record<string, unknown>)
          .flatMap(([keyItem, items]) => {
            if (isArray(items)) {
              return items.flatMap((item) => getEncode([keyItem, item], { ...config, default: undefined }));
            }
            return getEncode([keyItem, items], { ...config, default: undefined });
          })
          .map((item) => [nameParam, item]) as [string, string | undefined][];
      } else if (params) {
        const itemConfig = Object.fromEntries(
          Object.entries(params).map(([pKey, pConfig]) => [
            pKey,
            { ...pConfig, prefix: prefix + (pConfig.prefix ?? '') },
          ])
        );
        return valueToArray(itemConfig, value[key], defaultParams?.[key] ?? config.default, urlSearchParams);
      }
    }
    if (
      config.isArray &&
      value?.[key] &&
      isArray(value[key]) &&
      !(value[key] as unknown[]).length &&
      typeof (defaultParams?.[key] ?? config.default) !== 'undefined' &&
      ((defaultParams?.[key] ?? config.default) as unknown[]).length
    ) {
      return [[nameParam, removeValueChar]];
    }
    if (!value && config.required && typeof (defaultParams?.[key] ?? config.default) !== 'undefined') {
      return [[nameParam, removeValueChar]];
    }
    return (getEncode(value?.[key], { ...config, default: defaultParams?.[key] ?? config.default })?.map((item) => [
      nameParam,
      item,
    ]) ?? [[nameParam, undefined]]) as [string, string | undefined][];
  });
}

/**
 * encode object in URLSearchParams
 *
 * @param configParams - parse config
 * @param value - value for encode
 * @param defaultParams - default value, if equal param not write
 * @param {URLSearchParams} urlSearchParams - source URLSearchParams for change
 *
 * @return {URLSearchParams}
 */
export function encodeQueryParams<T extends Record<string, unknown>>(
  configParams: ConfigParams,
  value: T,
  defaultParams?: T,
  urlSearchParams?: URLSearchParams
): URLSearchParams {
  const nextParams = new URLSearchParams(urlSearchParams ?? window.location.search);
  const updMap: Record<string, boolean> = {};
  valueToArray(configParams, value, defaultParams, urlSearchParams).forEach(([key, v]) => {
    if (key === '') {
      return;
    }
    if (!updMap[key]) {
      nextParams.delete(key);
    }
    updMap[key] = true;
    if (typeof v !== 'undefined') {
      nextParams.append(key, v);
    }
  });
  return nextParams;
}

/**
 * decode object by URLSearchParams
 *
 * @param configParams - parse config
 * @param defaultParams default value if not find in url
 * @param urlSearchParams - source URLSearchParams for change
 *
 * @return - parse object
 */
export function decodeQueryParams<T extends Record<string, unknown>>(
  configParams: ConfigParams,
  defaultParams?: T,
  urlSearchParams: URLSearchParams = new URLSearchParams(window.location.search)
): T | null {
  try {
    const res = Object.entries(configParams).map(([key, config]) => {
      const { urlKey, prefix = '', prefixArray, params, fromEntries } = config;
      if (config.isArray && prefixArray && params) {
        const arr = [];
        for (let i = 0; i < maxPrefixArray; i++) {
          let prefixA = prefixArray(i);
          const itemConfig = Object.fromEntries(
            Object.entries(params).map(([pKey, pConfig]) => [
              pKey,
              { ...pConfig, prefix: prefixA + (pConfig.prefix ?? '') },
            ])
          );

          const item = decodeQueryParams(
            itemConfig,
            config.struct
              ? config.default?.[key]?.[i]
              : (defaultParams?.[key] as unknown[])?.[i] ?? config.default?.[key]?.[i],
            urlSearchParams
          );
          if (item) {
            if (Object.values(item).some((v) => v === removeValueChar)) {
              return [key, arr];
            }
            arr.push(item);
          } else {
            break;
          }
        }
        if (!arr.length) {
          const _default = defaultParams?.[key] ?? config.default;
          return [key, _default && deepClone(_default)];
        }
        return [key, arr];
      } else if (fromEntries) {
        const values = urlSearchParams.getAll(prefix + (urlKey ?? key));
        if (values.length === 0) {
          return [key, defaultParams?.[key] ?? config.default];
        }
        if (values.length === 1 && values[0] === removeValueChar) {
          return [key, {}];
        }

        const items = getDecode(values, { ...config, default: defaultParams?.[key] ?? config.default }) as (
          | [string, unknown]
          | undefined
        )[];
        return [
          key,
          items.reduce((res, item) => {
            if (item) {
              if (config.isArray) {
                res[item[0]] = res[item[0]] ?? [];
                (res[item[0]] as unknown[]).push(item[1]);
              } else if (!res.hasOwnProperty(item[0])) {
                //add only first
                res[item[0]] = item[1];
              }
            }
            return res;
          }, {} as Record<string, unknown | unknown[]>),
        ];
      } else if (params) {
        const itemConfig = Object.fromEntries(
          Object.entries(params).map(([pKey, pConfig]) => [
            pKey,
            { ...pConfig, prefix: prefix + (pConfig.prefix ?? '') },
          ])
        );
        return [key, decodeQueryParams(itemConfig, { ...(defaultParams?.[key] ?? {}) }, urlSearchParams)];
      }
      const values = urlSearchParams.getAll(prefix + (urlKey ?? key));

      if (config.required && values.length === 0 && !(defaultParams?.[key] ?? config.default)) {
        throw new Error('required param not find ' + prefix + (urlKey ?? key));
      }
      if (config.isArray && values.length === 1 && values[0] === removeValueChar) {
        return [key, []];
      }
      if (config.isArray && values.length === 0) {
        return [key, defaultParams?.[key] ?? config.default];
      }
      return [key, getDecode(values, { ...config, default: defaultParams?.[key] ?? config.default })];
    });
    return Object.fromEntries(res);
  } catch (e) {
    return null;
  }
}
