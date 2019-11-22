import React from 'react';
import { createBrowserHistory } from "history";
import { Router, Route, Switch, Redirect } from "react-router-dom";
import { Sidebar } from './Sidebar/Sidebar';
import { Releases } from './pages/Releases/Releases';
import { History } from './pages/History/History';
import { Pipelines } from './pages/Pipelines/Pipelines';

const hist = createBrowserHistory();

export const App = () => {
  return (
    <Router history={hist}>
      <Sidebar />
      <Switch>
        <Route path="/releases" component={Releases} />
        <Route path="/history" component={History} />
        <Route path="/pipelines" component={Pipelines} />
        <Redirect from="/" to="/releases" />
      </Switch>
    </Router>
  )
}