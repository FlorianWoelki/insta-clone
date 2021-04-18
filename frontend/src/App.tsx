import React from 'react';
import { Switch, Route, BrowserRouter as Router } from 'react-router-dom';
import DefaultLayout from './layout/DefaultLayout';
import Direct from './views/Direct';
import Explore from './views/Explore';
import Home from './views/Home';
import Messages from './views/Messages';
import Notifications from './views/Notifications';
import Settings from './views/Settings';
import Stats from './views/Stats';

function App() {
  return (
    <Router>
      <DefaultLayout>
        <Switch>
          <Route exact path="/">
            <Home />
          </Route>
          <Route path="/direct">
            <Direct></Direct>
          </Route>
          <Route path="/explore">
            <Explore></Explore>
          </Route>
          <Route path="/messages">
            <Messages></Messages>
          </Route>
          <Route path="/notifications">
            <Notifications></Notifications>
          </Route>
          <Route path="/settings">
            <Settings></Settings>
          </Route>
          <Route path="/stats">
            <Stats></Stats>
          </Route>
        </Switch>
      </DefaultLayout>
    </Router>
  );
}

export default App;
