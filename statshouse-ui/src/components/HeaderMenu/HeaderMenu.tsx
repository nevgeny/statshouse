// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import React, { useCallback } from 'react';
import { NavLink, useLocation } from 'react-router-dom';
import { ReactComponent as SVGLightning } from 'bootstrap-icons/icons/lightning.svg';
import { ReactComponent as SVGGridFill } from 'bootstrap-icons/icons/grid-fill.svg';
import { ReactComponent as SVGPlus } from 'bootstrap-icons/icons/plus.svg';
import { ReactComponent as SVGCardList } from 'bootstrap-icons/icons/card-list.svg';
import { ReactComponent as SVGBrightnessHighFill } from 'bootstrap-icons/icons/brightness-high-fill.svg';
import { ReactComponent as SVGMoonStarsFill } from 'bootstrap-icons/icons/moon-stars-fill.svg';
import { ReactComponent as SVGCircleHalf } from 'bootstrap-icons/icons/circle-half.svg';
import { ReactComponent as SVGLightbulbFill } from 'bootstrap-icons/icons/lightbulb-fill.svg';
// import { ReactComponent as SVGGear } from 'bootstrap-icons/icons/gear.svg';
import cn from 'classnames';
import produce from 'immer';

import { HeaderMenuItem } from './HeaderMenuItem';
import {
  selectorParams,
  selectorPlotList,
  selectorPromqltestfailed,
  selectorSetParams,
  selectorSetTheme,
  selectorThemeName,
  THEMES,
  useStore,
} from '../../store';
import { currentAccessInfo, logoutURL } from '../../common/access';
import { HeaderMenuItemPlot } from './HeaderMenuItemPlot';
import css from './style.module.css';
import { parseParamsFromUrl } from '../../common/plotQueryParams';

const themeIcon = {
  [THEMES.Light]: SVGBrightnessHighFill,
  [THEMES.Dark]: SVGMoonStarsFill,
  [THEMES.Auto]: SVGCircleHalf,
};

export type HeaderMenuProps = {
  className?: string;
};

export const HeaderMenu: React.FC<HeaderMenuProps> = ({ className }) => {
  const params = useStore(selectorParams);
  const setParams = useStore(selectorSetParams);
  const menuPlots = useStore(selectorPlotList);
  const location = useLocation();
  const ai = currentAccessInfo();

  const themeName = useStore(selectorThemeName);
  const setTheme = useStore(selectorSetTheme);

  const promqltestfailed = useStore(selectorPromqltestfailed);

  const isView = location.pathname === '/view';
  const isDashList = location.pathname === '/dash-list';
  // const isAdminGroup = location.pathname === '/admin/group';

  const onPasteClipboard = useCallback(() => {
    (navigator.clipboard.readText ? navigator.clipboard.readText() : Promise.reject())
      .then((url) => {
        const parseParams = parseParamsFromUrl(url);
        if (parseParams.plots.length) {
          setParams(
            produce((p) => {
              p.plots = [...p.plots, ...parseParams.plots];
            })
          );
        }
      })
      .catch(() => {
        const url = prompt('Paste url') ?? '';
        const parseParams = parseParamsFromUrl(url);
        if (parseParams.plots.length) {
          setParams(
            produce((p) => {
              p.plots = [...p.plots, ...parseParams.plots];
            })
          );
        }
      });
  }, [setParams]);

  const onDark = useCallback(() => {
    setTheme(THEMES.Dark);
  }, [setTheme]);
  const onLight = useCallback(() => {
    setTheme(THEMES.Light);
  }, [setTheme]);
  const onAuto = useCallback(() => {
    setTheme(THEMES.Auto);
  }, [setTheme]);

  return (
    <div className={cn('sticky-top align-self-start', css.navOuter, className)}>
      <ul className={`nav pb-5  d-flex flex-column ${css.nav}`}>
        <HeaderMenuItem icon={SVGLightning} title="Home" to="/view" description="StatsHouse">
          <li className={css.splitter}></li>
          <li className="nav-item">
            <NavLink className="nav-link" to="/admin/create" end>
              Create&nbsp;metric
            </NavLink>
          </li>
          <li className={css.splitter}></li>
          <li className="nav-item">
            <a
              className="nav-link"
              href="https://github.com/VKCOM/statshouse/discussions/categories/announcements"
              target="_blank"
              rel="noreferrer"
            >
              News
            </a>
          </li>
          <li className="nav-item">
            <a
              className="nav-link"
              href="https://github.com/VKCOM/statshouse#documentation"
              target="_blank"
              rel="noreferrer"
            >
              Documentation
            </a>
          </li>
          <li className="nav-item">
            <a className="nav-link" href="/openapi/" target="_blank">
              OpenAPI
            </a>
          </li>
          <li className="nav-item">
            <a
              className="nav-link"
              href="https://github.com/VKCOM/statshouse/discussions/categories/q-a"
              target="_blank"
              rel="noreferrer"
            >
              Support
            </a>
          </li>
          <li className="nav-item">
            <NavLink className="nav-link" to="/doc/faq" end>
              FAQ
            </NavLink>
          </li>
          {!!ai.user && (
            <>
              <li className={css.splitter}></li>
              <li className="nav-item">
                <span className="nav-link text-secondary">{ai.user}</span>
              </li>
              <li className="nav-item">
                <a className="nav-link" href={logoutURL} title="Log out">
                  Log&nbsp;out
                </a>
              </li>
            </>
          )}
        </HeaderMenuItem>
        <HeaderMenuItem icon={themeIcon[themeName] ?? SVGBrightnessHighFill} title="Theme">
          <li className={css.splitter}></li>
          <li className={cn('nav-item', themeName === THEMES.Dark && 'bg-primary-subtle')}>
            <span role="button" className="nav-link" title="Set dark theme" onClick={onDark}>
              Dark
            </span>
          </li>
          <li className={cn('nav-item', themeName === THEMES.Light && 'bg-primary-subtle')}>
            <span role="button" className="nav-link" title="Set light theme" onClick={onLight}>
              Light
            </span>
          </li>
          <li className={cn('nav-item', themeName === THEMES.Auto && 'bg-primary-subtle')}>
            <span role="button" className="nav-link" title="Set auto theme" onClick={onAuto}>
              Auto
            </span>
          </li>
        </HeaderMenuItem>
        {/*{ai.admin && (
          <HeaderMenuItem
            icon={SVGGear}
            to="/settings/prometheus"
            title="Setting"
            className={cn(isAdminGroup && css.activeItem)}
          ></HeaderMenuItem>
        )}*/}
        <HeaderMenuItem
          icon={SVGCardList}
          to="/dash-list"
          title="Dashboard list"
          className={cn(isDashList && css.activeItem)}
        ></HeaderMenuItem>
        <HeaderMenuItem
          icon={SVGGridFill}
          indexPlot={-1}
          title="Dashboard"
          className={cn(params.tabNum < 0 && isView && css.activeItem)}
        >
          {/*<li className={css.splitter}></li>
          <li className="nav-item">
            <PlotLink className="nav-link" indexPlot={-2} title="Dashboard Setting">
              Dashboard&nbsp;Setting
            </PlotLink>
          </li>*/}
        </HeaderMenuItem>
        {menuPlots.map((item) => (
          <HeaderMenuItemPlot key={item.indexPlot} indexPlot={item.indexPlot} />
        ))}
        <HeaderMenuItem icon={SVGPlus} indexPlot={params.plots.length} title="Duplicate plot to new tab">
          <li className={css.splitter}></li>
          <li className="nav-item">
            <span role="button" className="nav-link" title="Paste new tab from clipboard" onClick={onPasteClipboard}>
              Paste&nbsp;Clipboard
            </span>
          </li>
        </HeaderMenuItem>
        {ai.developer && promqltestfailed && (
          <HeaderMenuItem
            icon={SVGLightbulbFill}
            title="promqltestfailed"
            className={css.secondDanger}
          ></HeaderMenuItem>
        )}
      </ul>
    </div>
  );
};
