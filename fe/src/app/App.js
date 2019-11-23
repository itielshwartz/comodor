import React from 'react';
import { createBrowserHistory } from "history";
import { Router, Route, Switch, Redirect } from "react-router-dom";
import { Sidebar } from './Sidebar/Sidebar';
import { Releases } from './pages/Releases/Releases';
import { History } from './pages/History/History';
import { Pipelines } from './pages/Pipelines/Pipelines';
import { SingleRelease } from './pages/Releases/SingleRelease/SingleRelease';

const hist = createBrowserHistory();

export const App = () =>   (
  <Router history={hist}>
    <Sidebar />
      <main>
        <Switch>
          <Route path="/releases/:id" component={SingleRelease} />
          <Route path="/releases" component={Releases} />
          <Route path="/history" component={History} />
          <Route path="/pipelines" component={Pipelines} />
          <Redirect from="/" to="/releases" />
        </Switch>
      </main>
  </Router>
)
