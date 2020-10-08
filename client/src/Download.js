import React from 'react';

function Download({rows, page}) {
    return (
        <div className="Download">
            <button onClick={() => downloadPageData(page)}>{"Download Page Data"}</button>
            {'  '}
            <button onClick={() => downloadRowsData(rows)}>{"Download All Data"}</button>
        </div>
    )
}

function downloadPageData(page) {
    let csv = json2csv(page.map(data => data.original))
    downloadCSVFile(csv, "page_data.csv")
}

function downloadRowsData(rows) {
    let csv = json2csv(rows.map(data => data.original))
    downloadCSVFile(csv, "all_data.csv")
}

function downloadCSVFile(csv, filename) {
    let exportData = 'data:text/csv;charset=utf-8,';
    exportData += csv;
    let encodedURI = encodeURI(exportData);
    let link = document.createElement("a");
    link.setAttribute("href", encodedURI);
    link.setAttribute("download", filename);
    document.body.appendChild(link);
    link.click()
}

function json2csv(data) {
    const replacer = (key, value) => value === null ? '' : value // specify how you want to handle null values here
    const header = Object.keys(data[0])
    let csv = data.map(row => header.map(fieldName => JSON.stringify(row[fieldName], replacer)).join(','))
    csv.unshift(header.join(','))
    csv = csv.join('\r\n')
    return csv
}

export default Download
