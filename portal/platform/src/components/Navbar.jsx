import { Link } from 'react-router-dom';
import '../ctStyle.css';
import {useScrollPosition} from '../scripts/scrollPosition';

import logo from '../images/icons/grayScheme/CTGrayLogo.png';
import homeIcon from '../images/icons/grayScheme/CTGrayHomeIcon.png';
import labIcon from '../images/icons/grayScheme/CTGrayLabIcon.png';
import dataIcon from '../images/icons/grayScheme/CTGrayDataIcon.png';
import searchIcon from '../images/icons/grayScheme/CTGraySearchIcon.png';
import gearIcon from '../images/icons/grayScheme/CTGrayGearIcon.png';

function Navbar(active) {
    const scrollPosition = useScrollPosition();

    let passClass = "";
    if (scrollPosition > 0) {
        passClass += " scrolled";
    }

    return (
        <div>
            <nav id="navbar" className={passClass}>
                {active.active == "home" ? (
                    <div className="navLinkContainer active">
                        <Link to="/home" className="navLink mainLink">
                            <img src={homeIcon} alt="Home icon" className="logo" />
                            Home
                        </Link>
                    </div>
                ) : (
                    <div className="navLinkContainer">
                        <Link to="/home" className="navLink mainLink">
                            <img src={homeIcon} alt="Home icon" className="logo" />
                            Home
                        </Link>
                    </div>
                )}

                <div className='smallBottom'></div>

                {active.active == "projects" ? (
                    <div className="navLinkContainer active">
                        <Link to="/projects" className="navLink mainLink">
                            <img src={labIcon} alt="File icon" className="logo" />
                            Projects
                        </Link>
                    </div>
                ) : (
                    <div className="navLinkContainer">
                        <Link to="/projects" className="navLink mainLink">
                            <img src={labIcon} alt="File icon" className="logo" />
                            Projects
                        </Link>
                    </div>
                )}

                <div className='smallBottom'></div>

                {active.active == "reports" ? (
                    <div className="navLinkContainer active">
                        <Link to="/reports" className="navLink mainLink">
                            <img src={dataIcon} alt="File icon" className="logo" />
                            Reports
                        </Link>
                    </div>
                ) : (
                    <div className="navLinkContainer">
                        <Link to="/reports" className="navLink mainLink">
                            <img src={dataIcon} alt="File icon" className="logo" />
                            Reports
                        </Link>
                    </div>
                )}

                <div className='smallBottom'></div>

                {active.active == "search" ? (
                    <div className="navLinkContainer active">
                        <Link to="/search" className="navLink mainLink">
                            <img src={searchIcon} alt="Search icon" className="logo" />
                            Search
                        </Link>
                    </div>
                ) : (
                    <div className="navLinkContainer">
                        <Link to="/search" className="navLink mainLink">
                            <img src={searchIcon} alt="Search icon" className="logo" />
                            Search
                        </Link>
                    </div>
                )}

                {active.active == "settings" ? (
                    <div className="navLinkContainer autoTopMargin active">
                        <Link to="/settings" className="navLink mainLink">
                            <img src={gearIcon} alt="Gear icon" className="logo" />
                            Settings
                        </Link>
                    </div>
                ) : (
                    <div className="navLinkContainer autoTopMargin">
                        <Link to="/settings" className="navLink mainLink">
                            <img src={gearIcon} alt="Gear icon" className="logo" />
                            Settings
                        </Link>
                    </div>
                )}

                <div className='smallBottom'></div>

                {active.active == "account" ? (
                    <div className="navLinkContainer smallBottom active">
                        <Link to="/account" className="navLink mainLink">
                            <img src={logo} alt="ChakraTech logo" className="logo" />
                            Account
                        </Link>
                    </div>
                ) : (
                    <div className="navLinkContainer smallBottom">
                        <Link to="/account" className="navLink mainLink">
                            <img src={logo} alt="ChakraTech logo" className="logo" />
                            Account
                        </Link>
                    </div>
                )}
            </nav>
        </div>
    )
}

export default Navbar;