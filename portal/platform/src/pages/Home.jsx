//home page
import React, { useState } from "react";
import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
import SideBar from "../components/SideBar";

import FadeIn from "../scripts/fadeIn";

import handWave from "../images/pictures/CTHandWave.png";

import * as dataCaller from '../scripts/dataCaller';
import useUsers from "./UserData"

function Home() {
    //update dataCaller with server data
    dataCaller.setUserData(useUsers());

    let name;
    try {
        name = dataCaller.getFirstName(); //get from database
    } catch (err) {
        name = "Guest";
    }
    
    //State to track form input values
    const [title, setTitle] = useState('');
    const [log, setLog] = useState('');

    //Handle form focus/blur 
    function handleFocus(e) {
        e.target.parentNode.classList.add('focused');
    }
    function handleBlur(e) {
        e.target.parentNode.classList.remove('focused');
        if(e.target.value === '') {
            if (e.target.parentNode.classList.contains('has-content')) {
                e.target.parentNode.classList.remove('has-content');
            }
        } else {
            e.target.parentNode.classList.add('has-content');
        }
    }

    return (
        <div>
            <FadeIn time={250} />
            <Navbar active="home" />

            <main className="page">
                <div className="container">
                    <div className="containerSub fadeIn">
                        <h1>Welcome, {name}</h1>
                        <p className="mediumBottom">Let's get started today!</p>
                        <img className="sectionBG" src={handWave}></img>
                    </div>

                    <div className="containerSub fadeIn">
                        <h2>Quick Input</h2>
                        <form>
                            <div>
                                <label htmlFor="title">Report Title</label>
                                <input type="text" id="title" name="title" value={title} onFocus={handleFocus} onBlur={handleBlur} onChange={(e) => setTitle(e.target.value)}/>
                            </div>
                            <div>
                                <label htmlFor="log">Log text...</label>
                                <textarea id="log" name="log" rows="4" cols="50" value={log} onFocus={handleFocus} onBlur={handleBlur} onChange={(e) => setLog(e.target.value)}></textarea>
                            </div>
                            <div className="formButtonContainer">
                                <button type="submit">Submit</button>
                                <button type="button" className="extra">Attach Files</button>
                            </div>
                        </form>
                    </div>
                </div>

                <SideBar />
            </main>

            <Footer />
        </div>
    )
}

export default Home;
