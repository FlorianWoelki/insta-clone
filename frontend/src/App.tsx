import React from 'react';
import { Switch, Route, BrowserRouter as Router } from 'react-router-dom';
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
      <div className="App">
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
      </div>
    </Router>
  );
}

export default App;
