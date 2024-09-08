//reports list page

//import main components
import React, {useEffect, useState} from "react";
import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
import '../ctStyle.css';

//import extra components
import { Link } from 'react-router-dom';
import Popup from "../components/Popup";
import { createRoot } from 'react-dom/client';

//import scripts
import FadeIn from "../scripts/fadeIn";
import useReports from "./ReportData"

function Reports() {
    const reportList = useReports();

    //popup creation
    const createReport = () => { //adding event
        const popup = document.createElement('div');
        createRoot(popup).render(<Popup type="addReport" header="Add Report" />);
        document.querySelector('main').appendChild(popup);
    };

    return (
        <div>
            <FadeIn time={50} />
            <Navbar active="reports" />

            <main className="commonPage">
                <div className="fullContainer">
                    <div className="containerSub">
                        <div className="flexBetween fadeIn">
                            <h1 className="noBottom">Reports</h1>
                            <button onClick={createReport}>Create Report<div className="addIcon">+</div></button>
                        </div>
                        <p className="mediumBottom fadeIn">Write about your findings and share information.</p>

                        <div className="listArea fadeIn">
                        <div className="searchContainer fadeIn">
                                {/* Search Box */}
                                <input id="search" type="text" placeholder="Search..." />
                                <button>🔎︎</button>

                                {/* Dropdown Filters */}
                                <select id="filter1">
                                    <option value="lastModified">Last Modified</option>
                                    <option value="created">Created</option>
                                    <option value="author">Author</option>
                                    <option value="name">Name</option>
                                </select>
                                <select id="filter2">
                                    <option value="ascending">Ascending</option>
                                    <option value="descending">Descending</option>
                                </select>
                            </div>

                            <div className="listContainer fadeIn">
                                {/* Header */}
                                <div className="listHeader">
                                    <div className="headerItem">Name</div>
                                    <div className="headerItem">Author</div>
                                    <div className="headerItem">Date Created</div>
                                    <div className="headerItem">Date Modified</div>
                                    <div className="headerItem">Logs</div>
                                </div>

                                {/* Report List */}
                                {reportList.map((report, index) => (
                                    <Link to={`/reports/${report.id}`} className={`listRow ${index % 2 !== 0 ? 'oddRow' : ''} `} key={index}>
                                        <div className="listItem titleLI">{report.title}</div>
                                        <div className="listItem titleLI">{report.owner}</div>
                                        <div className="listItem">{report.date}</div>
                                        <div className="listItem">{report.lastModified}</div>
                                        <div className="listItem">{report.logs.length}</div>
                                    </Link>
                                ))}
                            </div>
                        </div>
                    </div>
                </div>
            </main>

            <Footer />
        </div>
    )
}

export default Reports;
