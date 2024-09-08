//fullscreen calendar page
import React from "react";

import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
import Calendar from "../components/Calendar";

import FadeIn from "../scripts/fadeIn";

function Home() {

    return (
        <div>
            <FadeIn time={200} />
            <Navbar />

            <main className="calendarPage">
                <Calendar size="large" />
            </main>

            <Footer />
        </div>
    )
}

export default Home;