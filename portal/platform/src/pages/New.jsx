import React, { useState } from "react";
import Navbar from "../components/Navbar";
import FadeIn from "../scripts/fadeIn";

function New() {
    //State to track form input values
    const [key1, setKey1] = useState('');
    const [key2, setKey2] = useState('');

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

    function handleSubmit(e) {
        e.preventDefault();
        //to be added with login
        console.log("Form submitted!");
    }

    return (
        <div>
            <FadeIn time={250} />
            <Navbar />

            <main className="page superCenter">
                <div className="loginContainer fadeIn">
                    <h1>Welcome</h1>
                    <p>Enter your credentials to get started.</p>

                    <form onSubmit={handleSubmit}>
                        <div>
                            <label htmlFor="email">Email Address</label>
                            <input type="text" id="email" name="email" autoComplete="email" value={key1} onFocus={handleFocus} onBlur={handleBlur} onChange={(e) => setKey1(e.target.value)}/>
                        </div>
                        <div>
                            <label htmlFor="password">Password</label>
                            <input type="password" id="password" name="password" autoComplete="email" value={key2} onFocus={handleFocus} onBlur={handleBlur} onChange={(e) => setKey2(e.target.value)}/>
                        </div>
                        <div className="formButtonContainer">
                            <button type="submit">Log In</button>
                            <button type="button" className="extra">Show Password</button>
                        </div>
                    </form>
                </div>
            </main>
        </div>
    )
}

export default New;