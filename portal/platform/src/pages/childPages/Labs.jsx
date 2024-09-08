//labs list page

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

//import assets
import labIcon from "../../images/icons/grayScheme/CTGrayLabIcon.png";

//import scripts
import FadeIn from "../../scripts/fadeIn";
import * as dataCaller from '../../scripts/dataCaller';

function Labs() {

    let id = window.location.pathname.split("/")[window.location.pathname.split("/").length - 1];

    let project = dataCaller.getProject(id); //get project object
    let labs = dataCaller.getProjectLabs(id); //get project labs

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
                    <div className="containerSub">
                        <div className="fadeIn">
                            <div className="flexBetween">
                                <div className="backButtonWrap">
                                    <Link to="/projects" className="backButton">{"<"} Projects</Link>
                                </div>
                                <div className="flexBetween flexGap">
                                    <button onClick={createLab}>Add Lab<div className="addIcon">+</div></button>
                                    <button onClick={editProject}>Edit Project<div className="editIcon">✎</div></button>
                                </div>
                            </div>
                            
                            <h1 className="noBottom smallTop">{project.title}</h1>
                            {project.description !== "" ? <p className="noBottom">{project.description}</p> : null}
                        </div>

                        <div className="line mediumTop mediumBottom fadeIn"></div>

                        <div className="fadeIn labArea">
                            {/*Display all labs*/}
                            {labs.map((labs, index) => (
                                <Link to={`/projects/${id}/${index}`} className="labContainer" key={index}>
                                    <div className="flexCenter fullHeight">
                                        {labs.thumbnail === "" ? (
                                            <img className="labImage default" src={labIcon}></img>
                                        ) : (
                                            <img className="labImage extraWidth" src={labs.thumbnail} alt={`Lab ${index+1} Thumbnail`}></img>
                                        )}

                                        <div className="fullWidth fullHeight higherIndex labInfoAdjust">
                                            <div className="labInfo">
                                                <p>{labs.title}</p>
                                                {labs.runs.length > 0 ? <p className="smallBottom smallText">Last Run {labs.runs[labs.runs.length-1].date}</p> : <p className="smallBottom smallText">No Runs</p>}
                                            </div>
                                        </div>
                                    </div>
                                </Link>
                            ))}
                        </div>
                    </div>
                </div>
            </main>

            <Footer />
        </div>
    )
}

export default Labs;