import React from 'react';
import { useTable, useSortBy } from 'react-table';

const columnHeaders = {
    Header: "Player Rushing Data",
    columns: [
        {
            Header: "Player",
            accessor: "Player",
        },
        {
            Header: "Team",
            accessor: "Team",
        },
        {
            Header: "Position",
            accessor: "Pos",
        },
        {
            Header: "Attempts",
            accessor: "Att",
        },
        {
            Header: "Avg Attempts/Game",
            accessor: "Att/G",
        },
        {
            Header: "Total Yards",
            accessor: "Yds",
        },
        {
            Header: "Avg Yards/Attempt",
            accessor: "Avg",
        },
        {
            Header: "Yards/Game",
            accessor: "Yds/G",
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
        },
        {
            Header: "First Downs %",
            accessor: "1st%",
        },
        {
            Header: "20+ Yards Each",
            accessor: "20+",
        },
        {
            Header: "40+ Yards Each",
            accessor: "40+",
        },
        {
            Header: "Fumbles",
            accessor: "FUM",
        },
    ]
}

function Table({ data }) { 
    const playerColumns = React.useMemo(() => [columnHeaders], []);
    debugger;

    // Use the state and functions returned from useTable to build your UI
    const {
        getTableProps,
        getTableBodyProps,
        headerGroups,
        rows,
        prepareRow,
    } = useTable({ playerColumns, data })

    // Render the UI for your table
    return (
        <table {...getTableProps()}>
            <thead>
                {headerGroups.map(headerGroup => (
                    <tr {...headerGroup.getHeaderGroupProps()}>
                        {headerGroup.headers.map(column => (
                            <th {...column.getHeaderProps()}>{column.render('Header')}</th>
                        ))}
                    </tr>
                ))}
            </thead>
            <tbody {...getTableBodyProps()}>
                {rows.map((row, i) => {
                    prepareRow(row)
                    return (
                        <tr {...row.getRowProps()}>
                            {row.cells.map(cell => {
                                return <td {...cell.getCellProps()}>{cell.render('Cell')}</td>
                            })}
                        </tr>
                    )
                })}
            </tbody>
        </table>
    )
}

export default Table;
