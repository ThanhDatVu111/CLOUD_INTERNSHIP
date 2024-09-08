import './universal.css';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';

import NoPage from "./pages/NoPage"; //404

import New from "./pages/New"; //login
import Home from "./pages/Home"; //main
import Projects from "./pages/Projects"; //projects
import Reports from "./pages/Reports"; //reports

import FullCalendar from "./pages/FullCalendar"; //full calendar

import Labs from "./pages/childPages/Labs"; //projects child
import Logs from "./pages/childPages/Logs"; //reports child

import Triggers from "./pages/childPages/grandchildPages/Triggers"; //projects grandchild

let isLoggedIn = localStorage.getItem('isLoggedIn');
let loggedID = localStorage.getItem('loggedID');

isLoggedIn = true; //change once login is implemented
loggedID = 0; //change once login is implemented

function App() {
  return (
    <div>
      <BrowserRouter>
        <Routes>

          {isLoggedIn ? (
            <>
              {/* Redirects */}
              <Route path="/" element={<Navigate to="/home" />} />
              <Route path="/login" element={<Navigate to="/home" />} />

              {/* Pages */}
              <Route path="/home" element={<Home />} />
              <Route path="/projects" element={<Projects />} />
                <Route path="/projects/:id" element={<Labs />} />
                  <Route path="/projects/:id/:id" element={<Triggers />} />
              <Route path="/reports" element={<Reports />} />
                <Route path="/reports/:id" element={<Logs />} />

              {/* Sub Pages */}
              <Route path="/calendar" element={<FullCalendar />} />
            </>
          ) : (
            <>
              {/* Login */}
              <Route path="*" element={<Navigate to="/login" />} />
              <Route path="/login" element={<New />} />
            </>
          )}
          {/* 404 */}
          <Route path="*" element={<NoPage />} />

        </Routes>
      </BrowserRouter>
    </div>
  );
}

export { isLoggedIn, loggedID };
export default App;