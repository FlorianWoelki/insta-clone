import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import Home from './views/Home';

function App() {
  return (
    <Router>
      <div className="App">
        <Home />
      </div>
    </Router>
  );
}

export default App;
