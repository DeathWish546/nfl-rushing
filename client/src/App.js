import React from 'react';
import logo from './logo.svg';
import './App.css';
import { playerData } from './rushing.js';
import Table from './Table';

function App() {
    return (
        <div className="App">
            <Table data={playerData} />
        </div>
    );
}

export default App;
