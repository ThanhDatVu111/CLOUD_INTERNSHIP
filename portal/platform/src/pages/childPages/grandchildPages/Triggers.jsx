//404 page
//pages that don't exist will be redirected here

//report detail view

//import main components
import React, {useEffect, useState} from "react";
import Navbar from "../../../components/Navbar";
import Footer from "../../../components/Footer";
import '../../../ctStyle.css';

//import extra components
import { Link } from 'react-router-dom';
import Popup from "../../../components/Popup";
import { createRoot } from 'react-dom/client';

//import scripts
import FadeIn from "../../../scripts/fadeIn";
import * as dataCaller from '../../../scripts/dataCaller';

function Triggers() {

    let id = window.location.pathname.split("/")[window.location.pathname.split("/").length - 2];
    let subId = window.location.pathname.split("/")[window.location.pathname.split("/").length - 1];

    let project = dataCaller.getProject(id); //get project object
    let labs = dataCaller.getProjectLabs(id); //get project labs
    let thisLab = labs[subId]; //get this lab
    let labRuns = [...thisLab.runs].reverse(); //runs in reverse order

    //popup creation
    const createLab = () => { //adding event
        const popup = document.createElement('div');
        createRoot(popup).render(<Popup type="addLog" header="Add Log" />);
        document.querySelector('main').appendChild(popup);
    }
    const editProject = () => { //adding event
        const popup = document.createElement('div');
        createRoot(popup).render(<Popup type="editReport" header="Edit Report" />);
        document.querySelector('main').appendChild(popup);
    }

    return (
        <div>
            <FadeIn time={50} />
            <Navbar active="projects" />

            <main className="commonPage">
                <div className="fullContainer">
                    <div className="containerSub fadeIn">
                        <div className="flexBetween">
                            <div className="flexBetween backButtonWrap">
                                <Link to="/projects" className="backButton">{"<"} Projects</Link>
                                <Link to={`/projects/${id}`} className="backButton">{"<"} {project.title}</Link>
                            </div>
                            <div className="flexBetween flexGap">
                                <button onClick={createLab}>Run<div className="editIcon">▶</div></button>
                                <button onClick={editProject}>Edit Lab<div className="editIcon">✎</div></button>
                            </div>
                        </div>
                        
                        <h1 className="noBottom smallTop">{thisLab.title}</h1>
                        {project.description !== "" ? <p className="noBottom">{thisLab.description}</p> : null}
                    </div>

                    {/* Current Lab Runs */}
                    {labRuns.length > 0 ? (
                        <div className="containerSub fadeIn">
                            <h2 className="noBottom smallTop">Latest Run</h2>

                            <div className="flexDown flexGap smallTop">
                                <div className="runRow flexBetween fadeIn">
                                    <div>
                                        <h4 className="noTop smallBottom">Triggered by {dataCaller.getName(labRuns[0].author)}</h4>
                                        <p className="noTop noBottom smallText">{dataCaller.getFirstName(labRuns[0].author)} triggered a run on {labRuns[0].date}</p>
                                    </div>
                                </div>

                                <div className="runRow flexBetween fadeIn">
                                    <div>
                                        <h4 className="noTop smallBottom">Results</h4>
                                        <p className="noTop noBottom">{labRuns[0].status.charAt(0).toUpperCase() + labRuns[0].status.slice(1)}</p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    ) : (
                        <div className="containerSub fadeIn">
                            <h2 className="noBottom smallTop">Latest Run</h2>
                            <p className="smallBottom">No runs available.</p>
                        </div>
                    )}

                    <div className="containerSub fadeIn">
                        <h2 className="noBottom smallTop">Previous Runs</h2>

                        {labRuns.length > 1 ? (
                            <div className="flexDown flexGap smallTop">
                                {/* Previous Lab Runs */}
                                {labRuns.map((run, index) => {
                                    if (index === 0) return null;
                                    return (
                                        <div className="flexDown flexGap" key={index}>
                                            <div className="runRow flexBetween fadeIn">
                                                <div>
                                                    <h4 className="noTop smallBottom">Triggered by {dataCaller.getName(run.author)}</h4>
                                                    <p className="noTop noBottom smallText">{dataCaller.getFirstName(run.author)} triggered a run on {run.date}</p>
                                                </div>
                                            </div>

                                            <div className="runRow flexBetween fadeIn">
                                                <div>
                                                    <h4 className="noTop smallBottom">Results</h4>
                                                    <p className="noTop noBottom">{run.status.charAt(0).toUpperCase() + run.status.slice(1)}</p>
                                                </div>
                                            </div>
                                        </div>
                                    )
                                })}
                            </div>
                        ) : (
                            <p className="smallBottom">No runs available.</p>
                        )}
                    </div>
                </div>
            </main>

            <Footer />
        </div>
    )
}

export default Triggers;