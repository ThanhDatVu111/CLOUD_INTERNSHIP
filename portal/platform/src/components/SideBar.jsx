import React, {useEffect, useState} from "react";
import { Link } from 'react-router-dom';

import Calendar from "./Calendar";

import * as dataCaller from '../scripts/dataCaller';

function SideBar() {
    const recentReports = dataCaller.getRecentReports(2);

    return (
        <div className="sideContainerHolder autoLeftMargin">
            <div className="sideContainer">
                <div className="sideBarSection smallBottom fadeIn">
                    <Calendar size="small" />
                </div>

                <div className="sideBarSection fadeIn">
                    <h2>Recent Reports</h2>
                    <div className="horizontalReports">
                        {recentReports.map((report, index) => (
                            <Link key={index} to={`/reports/${report.id}`} className="reportDisplay">
                                <img src={report.thumbnail}></img>
                                <h4>{report.title}</h4>
                                <p>{report.lastModified}</p>
                            </Link>
                        ))}
                    </div>
                </div>
            </div>
        </div>
    );
}

export default SideBar;