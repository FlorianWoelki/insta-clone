import React, { useState } from 'react';
import Avatar from './components/Avatar';
import Icon from './components/Icon';

function App() {
  return (
    <div className="App">
      <Avatar>
        <div className="absolute inset-0 flex items-center justify-center w-full h-full bg-blue-600 bg-opacity-50 rounded-full">
          <Icon name="plus" className="w-6 h-6 text-white" />
        </div>
      </Avatar>
      <Avatar></Avatar>
    </div>
  );
}

export default App;
