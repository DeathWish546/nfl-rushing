import React from 'react';
import axios from 'axios';
import './App.css';
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
    const apiUrl = "http://localhost:8080/players"

    const [error, setError] = React.useState(null);
    const [isLoaded, setIsLoaded] = React.useState(false)
    const [playerData, setPlayerData] = React.useState([])

    React.useEffect(() => {
        axios.get(apiUrl)
            .then((response) => {
                setIsLoaded(true);
                if (response.data.length && response.data.length) {
                    setPlayerData(response.data)
                }
            },
            (error) => {
                setIsLoaded(true);
                setError(error);
            })
    }, [])

    const playerColumns = React.useMemo(() => [columnHeaders], [])

    if (error){
        return <div>Error: {error.message}</div>;
    } else if (!isLoaded) {
        return (
            <div className="App">
                Loading Data...
            </div>
        );
    } else {
        return (
            <div className="App">
                <Table columns={playerColumns} data={playerData} />
            </div>
        );
    }
}

export default App;
