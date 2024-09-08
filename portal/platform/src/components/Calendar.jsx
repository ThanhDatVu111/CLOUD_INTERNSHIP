//import main components
import React, { useEffect, useRef } from 'react';
import { Link } from 'react-router-dom';
import '../ctStyle.css';

//import extra components
import Popup from "./Popup.jsx";
import { createRoot } from 'react-dom/client';

//import assets
import maximize from '../images/icons/grayScheme/CTGrayFullIcon.png';
import edit from '../images/icons/grayScheme/CTGrayEditIcon.png';
import trash from '../images/icons/grayScheme/CTGrayTrashIcon.png';

//import scripts
import * as dataCaller from '../scripts/dataCaller';

function Calendar({size}) {

    //to update the events list
    const eventsHeadRef = useRef(null); //today's events header
    const eventsListRef = useRef(null); //today's events
    const futureListRef = useRef(null); //upcoming events
    
    const takenDays = dataCaller.getPlannedDates();
    let todayEvents;
    let dateSelected; //date selected on the calendar
    
    //data for date display
    const today = new Date();
    const month = today.toLocaleString('default', { month: 'long' });
    const day = today.toLocaleString('default', { day: 'numeric'});
    const year = today.toLocaleString('default', { year: 'numeric'});

    //calendar creation
    const getDaysInMonth = (month, year) => {
        return new Date(year, month + 1, 0).getDate(); //create next month's date with day 0 (last day of prev month)
    };
    const createCalendarGrid = () => { //put calendar days into a grid
        const currentYear = today.getFullYear();
        const currentMonth = today.getMonth();
        const firstDay = new Date(currentYear, currentMonth, 1).getDay(); //weekday of the first day
        const daysInMonth = getDaysInMonth(currentMonth, currentYear);
        const days = [];
    
        //weekdays header
        days.push(
            <tr key="weekdays">
                <th>Sun</th>
                <th>Mon</th>
                <th>Tue</th>
                <th>Wed</th>
                <th>Thu</th>
                <th>Fri</th>
                <th>Sat</th>
            </tr>
        );
        let date = 1 - firstDay; //start from the correct day of the week
        for (let i = 0; i < 6; i++) { //max 6 rows for calendar
            const week = [];
            for (let j = 0; j < 7; j++) { //7 days in a week
                let className =
                    date < 1 || date > daysInMonth
                        ? "calendarDay gray" //not current month
                        : date === today.getDate()
                            ? "calendarDay today" //today
                            : "calendarDay white"; //current month
                
                let currentFullDate = (`${(currentMonth + 1).toString().padStart(2, '0')}/${date.toString().padStart(2, '0')}/${year}`);
                if (takenDays.includes(currentFullDate)) {
                    className += " green";
                }
    
                let displayDate = date;
                if (date < 1) {
                    // Previous month
                    const prevMonthDays = getDaysInMonth(currentMonth - 1, currentYear);
                    displayDate = prevMonthDays + date;
                } else if (date > daysInMonth) {
                    // Next month
                    displayDate = date - daysInMonth;
                }
    
                //autoselects today's date
                if (className.includes("today")) {
                    getDateData(displayDate);
                }

                //dates not part of the current month are grayed out
                if(size === "large") {
                    if(className.includes("gray")) {
                        week.push(
                            <td key={`${i}-${j}`} className={className}>
                                {displayDate}
                            </td>
                        );
                    //relevant dates
                    } else {
                        //get events for the day
                        todayEvents = dataCaller.getEvents(`${(today.getMonth() + 1).toString().padStart(2, '0')}/${displayDate.toString().padStart(2, '0')}/${today.getFullYear()}`, "nameTime");
                        
                        week.push(
                            <td key={`${i}-${j}`} className={className} onClick={() => getDateData(displayDate)}>
                                <div className='dateContent'>
                                    <div className='dateHead'>{displayDate}</div>
                                    {todayEvents.length > 0 ? ( //show events
                                        <div className='dayEventsContainer'>
                                            {todayEvents.map((event, index) => {
                                                //truncate long event names
                                                const maxLength = 35;
                                                let truncatedEvent = event;
                                                if (event.length > maxLength) {
                                                    truncatedEvent = event.substring(0, maxLength) + "...";
                                                }
        
                                                //return the event
                                                return (
                                                    <p key={index}>
                                                        {truncatedEvent}
                                                    </p>
                                                );
                                            })}
                                        </div>
                                    ) : (
                                        //no event
                                        <div className='dayEventsContainer'>
                                            <p>Empty</p>
                                        </div>
                                    )}
                                </div>
                            </td>
                        );
                    }
                } else {
                    if(className.includes("gray")) {
                        week.push(
                            <td key={`${i}-${j}`} className={className}>
                                {displayDate}
                            </td>
                        );
                    //today's date is highlighted
                    } else if (className.includes("today")){
                        week.push(
                            <td key={`${i}-${j}`} className={className} onClick={() => getDateData(displayDate)}>
                                {displayDate}
                            </td>
                        );
                    //
                    } else {
                        week.push(
                            <td key={`${i}-${j}`} className={className} onClick={() => getDateData(displayDate)}>
                                {displayDate}
                            </td>
                        );
                    }
                }
                date++;
            }
            if (week.length > 0) days.push(<tr className={`${size === 'large' ? 'fadeIn' : ''}`} key={i}>{week}</tr>);
        }
    
        return days;
    };

    //activated when a day is selected
    const getDateData = (selectedDate) => {
        const getterDate = `${(today.getMonth() + 1).toString().padStart(2, '0')}/${selectedDate.toString().padStart(2, '0')}/${today.getFullYear()}`;
        todayEvents = dataCaller.getEvents(getterDate, "name"); //event names
        const todayDates = dataCaller.getEvents(getterDate, "time"); //event times

        dateSelected = selectedDate; //making it more confusing than it needs to be
        
        //update the events list
        if (eventsHeadRef.current) { //update header
            {dateSelected == day ? (
                eventsHeadRef.current.textContent = "Today's Events"
            ) : (
                eventsHeadRef.current.textContent = "Events of day " + dateSelected
            )}
        }
        if (eventsListRef.current) { //update list
            todayEvents.length > 0 ? (
                eventsListRef.current.innerHTML = (
                    `<ul>
                        ${todayEvents.map((event, index) => {
                            const maxLength = 25;
                            if(event.length > maxLength) {
                                event = event.substring(0, maxLength) + "...";
                            }
                            return `<li class="quickIconParent" key=${index}>
                                <div class="quickIconContent">
                                    ${event}
                                    <div class="flexCenter">
                                        <button class="quickIcon editButton" data-eventName=${event.replace(/\s/g, "\\")} data-eventTime=${todayDates[index]}>
                                            <img src=${edit}>
                                        </button>
                                        <button class="quickIcon trashButton">
                                            <img src=${trash}>
                                        </button>
                                    </div>
                                </div>
                            </li>`;
                        }).join('')}
                    </ul>`
                )                
            ) : (
                eventsListRef.current.innerHTML = "<p>Nothing planned this day</p>"
            );

            //add event listeners to the edit and delete buttons
            let buttons = eventsListRef.current.querySelectorAll('.trashButton');
            buttons.forEach(button => {
                button.addEventListener('click', () => {
                    createBooleanPopup();
                });
            });
            buttons = eventsListRef.current.querySelectorAll('.editButton');
            buttons.forEach(button => {
                button.addEventListener('click', () => {
                    createEditPopup(button.getAttribute("data-eventName"), button.getAttribute("data-eventTime"));
                });
            });
        }
    };

    //activated to get the upcoming events after today
    const getFutureData = () => {
        const futureEvents = dataCaller.getFutureEvents(`${(today.getMonth() + 1).toString().padStart(2, '0')}/${dateSelected.toString().padStart(2, "0")}/${today.getFullYear()}`);
        if (futureListRef.current) { //update list
            futureEvents.length > 0 ? (
                futureListRef.current.innerHTML = (
                    `<ul>
                        ${futureEvents.map((event, index) => {
                            const maxLength = 20;
                            let thisDescription = event.description;
                            if(event.description.length > maxLength) {
                                thisDescription = event.description.substring(0, maxLength) + "...";
                            }
                            return `<li class="quickIconParent" key=${index}><div class="quickIconContent">${event.date}: ${thisDescription}<div class="flexCenter"><button class="quickIcon"><img src=${edit}></button><button class="quickIcon"><img src=${trash}></button></div></div></li>`
                        }).join('')}
                    </ul>`
                )
            ) : (
                futureListRef.current.innerHTML = "<p>Nothing upcoming</p>"
            );

            //add event listeners to the edit and delete buttons
            let buttons = futureListRef.current.querySelectorAll('.trashButton');
            buttons.forEach(button => {
                button.addEventListener('click', () => {
                    createBooleanPopup();
                });
            });
            buttons = futureListRef.current.querySelectorAll('.editButton');
            buttons.forEach(button => {
                button.addEventListener('click', () => {
                    createEditPopup(button.getAttribute("data-eventName"), button.getAttribute("data-eventTime"));
                });
            });
        }
    }

    //initial render
    useEffect(() => {
        const today = new Date();
        const selectedDate = today.getDate();
        getDateData(selectedDate);
        getFutureData(selectedDate);
    }, []);

    //popup creation
    const createTextPopup = () => { //adding event
        const popup = document.createElement('div');
        createRoot(popup).render(<Popup type="addEvent" header="Add Event" />);
        document.querySelector('main').appendChild(popup);
    }
    function createBooleanPopup() { //deleting event
        const popup = document.createElement('div');
        createRoot(popup).render(<Popup type="boolean" header="Delete Event" message="Are you sure you want to delete this event?" />);
        document.querySelector('main').appendChild(popup);
    }
    function createEditPopup(passPreMessage, passPreTime) { //editing event
        passPreMessage = passPreMessage.replace(/\\/g, " ");
        const popup = document.createElement('div');
        createRoot(popup).render(<Popup type="editEvent" header="Edit Event" preMessage={passPreMessage} preTime={passPreTime} />);
        document.querySelector('main').appendChild(popup);
    }

    if (size === 'small') {
        return (
            <div className="sideCalendar">
                <div className='sectionOptions'>
                    <Link to={"/calendar"}>
                        <img src={maximize}></img>
                    </Link>
                </div>
    
                <div>
                    <h2>{month} {day}, {year}</h2>
                    <table className="calendar">
                        <tbody>{createCalendarGrid()}</tbody>
                    </table>
                </div>
    
                <div id="todayEvents">
                    <h3 ref={eventsHeadRef} >Today's Events</h3>
                    <div ref={eventsListRef}>
                        <p>Nothing planned today</p>
                    </div>
                </div>
    
                <div>
                    <h3>Upcoming Events</h3>
                    <div ref={futureListRef}>
                        <p>Nothing upcoming</p>
                    </div>
                </div>
            </div>
        );
    } else {
        return (
            <div className="mainCalendar">
                <div className="calendarArea fadeIn">
                    <div className='flexCenter'>
                        <button>⮜</button>
                        <h2>{month} <span className='light'>{year}</span></h2>
                        <button>⮞</button>
                    </div>
                    <table className="calendar">
                        <tbody>{createCalendarGrid()}</tbody>
                    </table>
                </div>

                <div className='sideBarMini fadeIn'>
                    <div id="todayEvents">
                        <h3 ref={eventsHeadRef} >Today's Events</h3>
                        <div ref={eventsListRef}>
                            <p>Nothing planned today</p>
                        </div>
                        <button onClick={() => createTextPopup()} >Add Event<div className="addIcon">+</div></button>
                    </div>

                    <div>
                        <h3>Upcoming Events</h3>
                        <div ref={futureListRef}>
                            <p>Nothing upcoming</p>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default Calendar;