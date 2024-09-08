//project list page

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
import * as dataCaller from '../scripts/dataCaller';
import useProjects from "./ProjectData";

function Projects() {
    const projectList = useProjects()
    

    //popup creation
    const createProject = () => { //adding event
        const popup = document.createElement('div');
        createRoot(popup).render(<Popup type="addProject" header="Add Project" />);
        document.querySelector('main').appendChild(popup);
    }

    return (
        <div>
            <FadeIn time={50} />
            <Navbar active="projects" />

            <main className="commonPage">
                <div className="fullContainer">
                    <div className="containerSub">
                        <div className="flexBetween fadeIn">
                            <h1 className="noBottom">Projects</h1>
                            <button onClick={createProject}>Create Project<div className="addIcon">+</div></button>
                        </div>
                        <p className="mediumBottom fadeIn">Manage lab experiments and run tests.</p>

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
                                    <div className="headerItem">Labs</div>
                                </div>

                                {/* Project List */}
                                {projectList.map((project, index) => (
                                    <Link to={`/projects/${project.id}`} className={`listRow ${index % 2 !== 0 ? 'oddRow' : ''} `} key={index}>
                                        <div className="listItem titleLI">{project.title}</div>
                                        <div className="listItem authorLI">{project.owner}</div>
                                        <div className="listItem">{project.date}</div>
                                        <div className="listItem">{project.lastModified}</div>
                                        <div className="listItem">{project.labs.length}</div>
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

export default Projects;
