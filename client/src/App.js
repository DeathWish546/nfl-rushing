import React from 'react';
import logo from './logo.svg';
import './App.css';
import { playerData } from './rushing.js'; //TODO REMOVE
import Table from './Table';

const columnHeaders = {
    Header: "Player Rushing Data",
    columns: [
        {
            Header: "Player",
            accessor: "Player",
            disableSortBy: true,
        },
        {
            Header: "Team",
            accessor: "Team",
            disableSortBy: true,
        },
        {
            Header: "Position",
            accessor: "Pos",
            disableSortBy: true,
        },
        {
            Header: "Attempts",
            accessor: "Att",
            disableSortBy: true,
        },
        {
            Header: "Avg Attempts/Game",
            accessor: "Att/G",
            disableSortBy: true,
        },
        {
            Header: "Total Yards",
            accessor: "Yds",
        },
        {
            Header: "Avg Yards/Attempt",
            accessor: "Avg",
            disableSortBy: true,
        },
        {
            Header: "Yards/Game",
            accessor: "Yds/G",
            disableSortBy: true,
        },
        {
            Header: "Touchdowns",
            accessor: "TD",
        },
        {
            Header: "Longest Rush",
            accessor: "Lng",
        },
        {
            Header: "First Downs",
            accessor: "1st",
            disableSortBy: true,
        },
        {
            Header: "First Downs %",
            accessor: "1st%",
            disableSortBy: true,
        },
        {
            Header: "20+ Yards Each",
            accessor: "20+",
            disableSortBy: true,
        },
        {
            Header: "40+ Yards Each",
            accessor: "40+",
            disableSortBy: true,
        },
        {
            Header: "Fumbles",
            accessor: "FUM",
            disableSortBy: true,
        },
    ]
}

function App() {
    const playerColumns = React.useMemo(() => [columnHeaders], [])
	const allData = React.useMemo(() => playerData, [])
    return (
        <div className="App">
            <Table columns={playerColumns} data={allData} />
        </div>
    );
}

export default App;
