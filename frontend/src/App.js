import React from 'react';
import './App.css';
import DataDisplay from './components/DataDisplay';

function App() {
  return (
      <div className="App">
        <header className="App-header">
          <h1>Главная страница</h1>
          <DataDisplay />
        </header>
      </div>
  );
}

export default App;
