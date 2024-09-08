import React, { useEffect, useState } from "react";

export function useReports() {
    const [reportList, setReportList] = useState([]);

    useEffect(() => {
        fetch("http://localhost:5174/reports")
            .then(res => res.json())
            .then(data => {
                setReportList(data);
            })
            .catch(err => console.error("Error getting reports", err));
    }, []);

    return reportList;
}

export default useReports;