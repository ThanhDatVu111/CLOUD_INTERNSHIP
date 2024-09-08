//import temporary data
import tempUserData from '../data/sampleUsers.json';
import tempProjectData from '../data/sampleProjects.json';
import tempReportData from '../data/sampleReports.json';

//temporary data to be updated by server
let userData = tempUserData;
let projectData = tempProjectData;
let reportData = tempReportData;

//login status and ID
import { isLoggedIn, loggedID } from "../App";

//SERVER FUNCTIONS-----------------------------------------------------

//sets user data to server-provided data
export function setUserData(data) {
    if (data != null && data != undefined && data != [] && data != {} && data.length > 0) {
        userData = data;
    }
    
}
//sets project data to server-provided data
export function setProjectData(data) {
    if (data != null && data != undefined && data != [] && data != {} && data.length > 0) {
        projectData = data;
    }
    
}
//sets report data to server-provided data
export function setReportData(data) {
    if (data != null && data != undefined && data != [] && data != {} && data.length > 0) {
        reportData = data;
    }
    
}

//returns all user data
export function getUserData() {
    return userData;
}
//returns all project data
export function getProjectData() {
    return projectData;
}
//returns all report data
export function getReportData() {
    return reportData;
}



//USER FUNCTIONS-------------------------------------------------------

//returns user as an object
export function getUser(id) {
    return userData[id];
}

//returns the user's name (optional id parameter)
export function getName(id) {
    if (id) {
        return userData[id].name;
    } else {
        return userData[loggedID].name;
    }
}
//returns the user's first name (optional id parameter)
export function getFirstName(id) {
    if (id) {
        return userData[id].name.split(' ')[0];
    } else {
        return userData[loggedID].name.split(' ')[0];
    }
}

//returns planned event dates for the user
export function getPlannedDates() {
    let plannedDates = [];
    for (let i = 0; i < userData[loggedID].calendar.length; i++) {
        plannedDates.push(userData[loggedID].calendar[i].date);
    }
    return plannedDates;
}
//returns planned events for the user on a given date
export function getEvents(date, type) {
    let events = [];
    for (let i = 0; i < userData[loggedID].calendar.length; i++) {
        if (userData[loggedID].calendar[i].date === date) {
            if (type === "name") {
                events.push(userData[loggedID].calendar[i].description);
            } else if (type === "time") {
                events.push(userData[loggedID].calendar[i].time);
            } else if (type === "date") {
                events.push(userData[loggedID].calendar[i].date);
            } else if (type === "nameTime") {
                let toPush = userData[loggedID].calendar[i].description;
                if (userData[loggedID].calendar[i].time != null) {
                    toPush += " at " + formatTime(userData[loggedID].calendar[i].time);
                }
                events.push(toPush);
            }
        }
    }
    return events;
}
//returns the rest of the events coming this month after the given date
export function getFutureEvents(date) {
    let events = [];
    for (let i = 0; i <userData[loggedID].calendar.length; i++) {
        if (userData[loggedID].calendar[i].date > date) {
            events.push(userData[loggedID].calendar[i]);
        }
    }
    return events;
}



//REPORT FUNCTIONS-----------------------------------------------------

//returns report as an object
export function getReport(id) {
    return reportData[id];
}

//returns array of report logs
export function getReportLogs(id) {
    return reportData[id].logs;
}

//returns report name
export function getReportName(id) {
    return reportData[id].title;
}

//refreshes reports and contributed reports for all users
export function refreshReports() {
    console.log("refreshing reports");
}

//refreshes the user's owned and contributed reports
export function refreshUserReports(user) {
    //get all reports and dates
    let myReports = [];
    let myOwnedReports = [];
    let myReportDates = [];
    for (let i = 0; i < reportData.length; i++) {
        if (reportData[i].owner.includes(user) || reportData[i].contributors.includes(user)) {
            myReports.push(reportData[i].id);
            myReportDates.push(reportData[i].lastModified);
            if (reportData[i].owner.includes(user)) {
                myOwnedReports.push(reportData[i].id);
            }
        }
    }

    //sort myReports based on date modified
    myReports.sort((a, b) => {
        const indexA = myReports.indexOf(a);
        const indexB = myReports.indexOf(b);
        const dateA = new Date(myReportDates[indexA]);
        const dateB = new Date(myReportDates[indexB]);
        return dateA.getTime() - dateB.getTime();
    });
    myReports.reverse();

    //update database with refreshed myReports and myOwnedReports
    userData[user].contributedReports = myReports;
    userData[user].ownedReports = myOwnedReports;

    return myReports;
}

//returns essential contributed report details
export function getRecentReports(amount) { //amount is the max number of reports to return
    refreshUserReports(loggedID); //ensures higher accuracy
    let recentReports = [];
    for (let i = 0; i < userData[loggedID].contributedReports.length && i < amount; i++) {
        const currentTargetReport = userData[loggedID].contributedReports[i];
        //title, lastModified, and thumbnail
        recentReports.push({
            id: reportData[currentTargetReport].id,
            title: reportData[currentTargetReport].title,
            lastModified: reportData[currentTargetReport].lastModified,
            thumbnail: reportData[currentTargetReport].thumbnail
        });
    }
    return recentReports;
}

//returns contributed and owned reports with more details
export function getRecentReportsPlus(amount) { //amount is the max number of reports to return
    refreshUserReports(loggedID); //ensures higher accuracy
    let recentReports = [];
    for (let i = 0; i < userData[loggedID].contributedReports.length && i < amount; i++) {
        const currentTargetReport = userData[loggedID].contributedReports[i];
        //all direct report details
        recentReports.push({
            id: reportData[currentTargetReport].id,
            title: reportData[currentTargetReport].title,
            description: reportData[currentTargetReport].description,
            date: reportData[currentTargetReport].date,
            lastModified: reportData[currentTargetReport].lastModified,
            lastModifiedBy: reportData[currentTargetReport].lastModifiedBy,
            thumbnail: reportData[currentTargetReport].thumbnail,
            owner: reportData[currentTargetReport].owner,
            contributors: reportData[currentTargetReport].contributors,
            logs: reportData[currentTargetReport].logs.length
        });
    }
    return recentReports;
}


//PROJECT FUNCTIONS----------------------------------------------------

//returns project as an object
export function getProject(id) {
    return projectData[id];
}

//returns array of project labs
export function getProjectLabs(id) {
    return projectData[id].labs;
}

//returns project name
export function getProjectName(id) {
    return projectData[id].title;
}

//refreshes projects and contributed projects for all users
export function refreshProjects() {
    console.log("refreshing projects");
}

//refreshes user's owned and contributed projects
export function refreshUserProjects(user) {
    //get all projects and dates
    let myProjects = [];
    let myOwnedProjects = [];
    let myProjectDates = [];
    for (let i = 0; i < projectData.length; i++) {
        if (projectData[i].owner.includes(user) || projectData[i].contributors.includes(user)) {
            myProjects.push(projectData[i].id);
            myProjectDates.push(projectData[i].lastModified);
            if (projectData[i].owner.includes(user)) {
                myOwnedProjects.push(projectData[i].id);
            }
        }
    }

    //sort myProjects based on date modified
    myProjects.sort((a, b) => {
        const indexA = myProjects.indexOf(a);
        const indexB = myProjects.indexOf(b);
        const dateA = new Date(myProjectDates[indexA]);
        const dateB = new Date(myProjectDates[indexB]);
        return dateA.getTime() - dateB.getTime();
    });
    myProjects.reverse();

    //update database with refreshed myProjects and myOwnedProjects
    userData[user].contributedProjects = myProjects;
    userData[user].ownedProjects = myOwnedProjects;

    return myProjects;
}

//returns contributed and owned projects with more details
export function getRecentProjectsPlus(amount) { //amount is the max number of projects to return
    refreshUserProjects(loggedID); //ensures higher accuracy
    let recentProjects = [];
    for (let i = 0; i < userData[loggedID].contributedProjects.length && i < amount; i++) {
        const currentTargetProject = userData[loggedID].contributedProjects[i];
        //all direct project details
        recentProjects.push({
            id: projectData[currentTargetProject].id,
            title: projectData[currentTargetProject].title,
            description: projectData[currentTargetProject].description,
            date: projectData[currentTargetProject].date,
            lastModified: projectData[currentTargetProject].lastModified,
            lastModifiedBy: projectData[currentTargetProject].lastModifiedBy,
            thumbnail: projectData[currentTargetProject].thumbnail,
            owner: projectData[currentTargetProject].owner,
            contributors: projectData[currentTargetProject].contributors,
            labs: projectData[currentTargetProject].labs.length
        });
    }
    return recentProjects;
}


//GENERAL FUNCTIONS----------------------------------------------------
//returns the time in a readable format
export function formatTime(time) {
    let timeArray = time.split(':');
    let hour = parseInt(timeArray[0]);
    let minute = timeArray[1];
    let suffix = "AM";
    if (hour > 12) {
        hour -= 12;
        suffix = "PM";
    }
    return hour + ":" + minute + " " + suffix;
}