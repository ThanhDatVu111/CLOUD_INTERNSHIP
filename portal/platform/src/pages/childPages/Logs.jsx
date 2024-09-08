//logs list page

//import main components
import React, {useEffect, useState} from "react";
import Navbar from "../../components/Navbar";
import Footer from "../../components/Footer";
import '../../ctStyle.css';

//import extra components
import { Link } from 'react-router-dom';
import Popup from "../../components/Popup";
import { createRoot } from 'react-dom/client';
import Profile from "../../components/Profile";
import FileView from "../../components/FileView";

//import scripts
import FadeIn from "../../scripts/fadeIn";
import * as dataCaller from '../../scripts/dataCaller';

function Logs() {
    let id = window.location.pathname.split("/")[window.location.pathname.split("/").length - 1];

    let report = dataCaller.getReport(id); //get report object
    let logs = dataCaller.getReportLogs(id); //get report logs

    //scroll to the bottom smoothly after 1 second
    setTimeout(() => {
        window.scrollTo({
            top: document.body.scrollHeight,
            behavior: "smooth"
        });
    }, 100);

    //log expanding
    function handleExpand(index) {
        document.getElementsByClassName("expand")[index].classList.toggle("closed"); //toggle expand
        document.getElementsByClassName("expandButton")[index].classList.toggle("rotate180"); //rotate button
    }

    //popup creation
    const createLog = () => { //adding event
        const popup = document.createElement('div');
        createRoot(popup).render(<Popup type="addLog" header="Add Log" />);
        document.querySelector('main').appendChild(popup);
    }
    const editReport = () => { //adding event
        const popup = document.createElement('div');
        createRoot(popup).render(<Popup type="editReport" header="Edit Report" />);
        document.querySelector('main').appendChild(popup);
    }
    const editLog = () => { //adding event
        const popup = document.createElement('div');
        createRoot(popup).render(<Popup type="editLog" header="Edit Log" />);
        document.querySelector('main').appendChild(popup);
    }
    const deleteLog = () => { //adding event
        const popup = document.createElement('div');
        createRoot(popup).render(<Popup type="deleteLog" header="Delete Log" />);
        document.querySelector('main').appendChild(popup);
    }

    return (
        <div>
            <FadeIn time={50} />
            <Navbar active="reports" />

            <main className="commonPage flexDown">
                {/* Header */}
                <div className="reportHead fullContainer higherIndex">
                    <div className="containerSub shadow">
                        <img src={report.thumbnail} />

                        <div className="flexBetween smallBottom">
                        <div className="backButtonWrap">
                                    <Link to="/reports" className="backButton">{"<"} Projects</Link>
                                </div>
                            <button onClick={editReport}>Edit Report<div className="editIcon">✎</div></button>
                        </div>
                        <h1>{report.title}</h1>
                        <p>{report.description}</p>
                    </div>
                </div>

                {/* Logs */}
                <div className="logArea">
                    <div className="logHolder"></div>
                    {logs.map((log, index) => (
                        <div className="fullContainer" key={index}>
                            <div className="reportLog containerSub fadeIn">

                                {/* Log Head */}
                                <div className="flexBetween">
                                    <h2>{log.title}</h2>
                                    <button onClick={() => handleExpand(index)} className="expandButton">▼</button>
                                </div>
                                <p className="logItem">{log.description}</p>

                                {/* Log Date/Tags */}
                                <div className="flexRight flexGap smallBottom">
                                    <div><i>{log.date}</i></div>
                                    {log.tags.map((tag, index) => (
                                        <div key={index} className="tag">{tag}</div>
                                    ))}
                                </div>

                                {/* Log Extras */}
                                <div className="expand closed">

                                    {/* Log Markdown */}
                                    <div className="mediumBottom">
                                        <h3>Markdown</h3>
                                        <Link>Read "{log.comment}"</Link>
                                    </div>

                                    {/* Log Attatchments */}
                                    <div className="mediumBottom">
                                        <h3>Attatchments</h3>
                                        {log.files.length > 0 ? (
                                            <div className="flexRight flexGap">
                                                {log.files.map((attachment, index) => {
                                                    const fileExtension = attachment.split('.').pop();
                                                    return (
                                                        <div key={index} className="flexRight">
                                                            <FileView type={fileExtension} src={attachment} title={"File " + (index+1)} description={fileExtension.toUpperCase() + " File"} />
                                                        </div>
                                                    );
                                                })}
                                            </div>
                                        ) : (
                                            <p>No Attachments</p>
                                        )}
                                    </div>

                                    {/* Log Dates */}
                                    <div className="flexRight flexGap">
                                        <div className="flexCenter flexDown flexGap">Created by <Profile user={log.contributors[0]} size={25} /><div className="logItem">{log.date}</div></div>
                                        <div className="flexCenter flexDown flexGap">Modified by <Profile user={log.lastModifiedBy} size={25} /><div className="logItem">{log.lastModified}</div></div>
                                    </div>
                                </div>

                                {/* Log Buttons */}
                                <div className="flexRight flexGap">
                                    <button onClick={editLog}>Edit<div className="iconSmall">✎</div></button>
                                    <button onClick={deleteLog}>Delete<div className="iconSmall">✖</div></button>
                                </div>

                            </div>
                        </div>
                    ))}

                    {/* New Log */}
                    <div className="fullContainer centerFlex">
                        <button onClick={createLog} className="addLog fadeIn">
                            <h2>Add New Log</h2>
                        </button>
                    </div>

                </div>
            </main>

            <Footer />
        </div>
    )
}

export default Logs;