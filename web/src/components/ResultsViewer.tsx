import { useMemo } from "react";
import { Column, useTable, useSortBy } from "react-table";

export interface WordCounts {
  word: string;
  count: number;
}

const ResultsViewer = ({ data }: { data: WordCounts[] }) => {
  const columns = useMemo<Column<WordCounts>[]>(
    () => [
      { Header: "Word", accessor: "word" },
      { Header: "Occurrences", accessor: "count" },
    ],
    []
  );

  const sortBy = useMemo(() => [{ id: "count", desc: true }], []);

  const { getTableProps, getTableBodyProps, headerGroups, rows, prepareRow } =
    useTable(
      {
        data,
        columns,
        initialState: {
          // @ts-ignore
          sortBy,
        },
      },
      useSortBy
    );

  const rowsSlice = rows.slice(0, 100) // TODO optimise, pagination etc instead

  return (
    <table {...getTableProps()}>
      <thead>
        {headerGroups.map((headerGroup) => (
          <tr {...headerGroup.getHeaderGroupProps()}>
            {headerGroup.headers.map((column: any) => (
              <th {...column.getHeaderProps(column.getSortByToggleProps())}>
                {column.render("Header")}
                <span>
                  {column.isSorted ? (column.isSortedDesc ? " 🔽" : " 🔼") : ""}
                </span>
              </th>
            ))}
          </tr>
        ))}
      </thead>
      <tbody {...getTableBodyProps()}>
        {rowsSlice.map((row, i) => {
          prepareRow(row);
          return (
            <tr {...row.getRowProps()}>
              {row.cells.map((cell) => {
                return <td {...cell.getCellProps()}>{cell.render("Cell")}</td>;
              })}
            </tr>
          );
        })}
      </tbody>
    </table>
  );
};

export default ResultsViewer;
