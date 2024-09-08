//404 page
//pages that don't exist will be redirected here

import Navbar from "../components/Navbar";
import Footer from "../components/Footer";
import FadeIn from "../scripts/fadeIn";

import { Link } from 'react-router-dom';

function NoPage() {

    return (
        <div>
            <FadeIn time={250} />
            <Navbar />

            <div className="noPage">
                <div className="bottomMargin3"></div>

                <h1 className="fadeIn">404 Error</h1>
                <h3 className="fadeIn">We couldn't find that page.</h3>

                <div className="bottomMargin3"></div>

                <div>
                    <p className="fadeIn">The page you're looking for may have been moved, deleted, or never existed. Please check the URL.</p>
                    <Link to="/" className="fadeIn">Go Home</Link>
                </div>
            </div>

            <Footer />
        </div>
    )
}

export default NoPage;