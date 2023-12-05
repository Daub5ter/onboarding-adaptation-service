// eslint-disable-next-line no-unused-vars
import React, {useEffect, useState} from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { Routes } from 'react-router-dom';
import Header from './pages/templates/Header.jsx';
import Footer from './pages/templates/Footer.jsx';
import Home from './pages/Home.jsx';
import Onboarding from './pages/Onboarding.jsx';
import Adapting from './pages/Adapting.jsx';
import Login from './pages/Login.jsx';
import './App.css';
import LoadSession from "./Auth/LoadSession";

function App() {
    const [isLoaded, setLoad] = useState(false);
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [email, setEmail] = useState('');

    useEffect(() => {
        const sessionToken = localStorage.getItem("session_token");
        if (sessionToken !== null) {
            LoadSession(sessionToken, setLoad, setIsLoggedIn, setEmail);
        } else {
            setLoad(true);
        }
    }, []);

    return (
        <>
            {isLoaded ?
            <Router>
                <div style={{ display: 'grid', gridTemplateRows: 'auto 1fr auto', minHeight: '100vh' }}>
                    <Header isLoggedIn={isLoggedIn} email={email} />
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/onboarding" element={<Onboarding isLoggedIn={isLoggedIn} email={email} isLoaded={isLoaded}/>} />
                        <Route path="/adaptation" element={<Adapting isLoggedIn={isLoggedIn} email={email} isLoaded={isLoaded}/>} />
                        <Route path="/login" element={<Login setIsLoggedIn={setIsLoggedIn} setEmail={setEmail} />} />
                    </Routes>
                    <Footer />
                </div>
            </Router>
                : <></>}
        </>
    );
}

export default App;
