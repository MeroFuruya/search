import {type RouteDefinition} from '@solidjs/router'
import { lazy } from 'solid-js'

export const routes: RouteDefinition[] = [
  {
    path: "/",
    component: lazy(() => import("./pages/Root")),
  },
  {
    path: "/settings",
    component: lazy(() => import("./pages/Settings")),
  }
]