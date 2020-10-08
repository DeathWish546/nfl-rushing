import React from 'react';
import { useTable, useSortBy, useFilters, usePagination } from 'react-table';
import styled from 'styled-components';
import DefaultColumnFilter from './Filter';
import Pagination from './Pagination';
import Download from './Download';
									//TODO add default sorting symbol to columns

const Styles = styled.div`
  padding: 1rem;

  table {
    border-spacing: 0;
    border: 1px solid black;

    tr {
      :last-child {
        td {
          border-bottom: 0;
        }
      }
    }

    th,
    td {
      margin: 0;
      padding: 0.5rem;
      border-bottom: 1px solid black;
      border-right: 1px solid black;

      :last-child {
        border-right: 0;
      }
    }
  }

  .pagination {
	padding: 0.5rem;
  }
`

function Table({ columns, data }) {
	const defaultColumn = React.useMemo(
		() => ({
			Filter: DefaultColumnFilter
		}),
		[]
	);

    const {
        getTableProps,
        getTableBodyProps,
		headerGroups,
		rows,
		prepareRow,
		page,
		canPreviousPage,
		canNextPage,
		pageOptions,
		pageCount,
		gotoPage,
		nextPage,
		previousPage,
		setPageSize,
		state: { pageIndex, pageSize },
    } = useTable(
        {
            columns,
            data,
			defaultColumn,
            disableMultiSort: true,
        },
        useFilters,
        useSortBy,
		usePagination,
    )

    // Render the UI for table
    return (
        <Styles>
            <table {...getTableProps()}>
                <thead>
                    {headerGroups.map(headerGroup => (
                        <tr {...headerGroup.getHeaderGroupProps()}>
                            {headerGroup.headers.map(column => (
                                <th {...column.getHeaderProps(column.getSortByToggleProps())}>
                                    {column.render('Header')}
									<div>
                                        {column.Header === "Player" ? column.render("Filter") : null}
                                        {column.canSort
                                            ? column.isSorted
                                                ? column.isSortedDesc
                                                    ? ' ⬇'
                                                    : ' ⬆'
                                                : ' ⇅'
                                            : ''
                                        }
                                    </div>
                                </th>
                            ))}
                        </tr>
                    ))}
                </thead>
                <tbody {...getTableBodyProps()}>
					{page.map((row, i) => {
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
			<Pagination
    			gotoPage={gotoPage}
    			canPreviousPage={canPreviousPage}
    			canNextPage={canNextPage}
    			previousPage={previousPage}
    			nextPage={nextPage}
    			pageCount={pageCount}
    			pageIndex={pageIndex}
    			pageSize={pageSize}
    			pageOptions={pageOptions}
    			setPageSize={setPageSize}
			/>
			<Download
				rows={rows}
				page={page}
			/>
        </Styles>
    )
};

export default Table;
