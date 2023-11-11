// eslint-disable-next-line no-unused-vars
import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { Routes } from 'react-router-dom';
import Header from './pages/templates/Header.jsx';
import Footer from './pages/templates/Footer.jsx';
import Home from './pages/Home.jsx';
import Onboarding from './pages/Onboarding.jsx';
import Adapting from './pages/Adapting.jsx';
import Login from './pages/Login.jsx';
import './App.css';


function App() {
    return (
        <Router>
            <div style={{ display: 'grid', gridTemplateRows: 'auto 1fr auto', minHeight: '100vh' }}>
                <Header />
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/onboarding" element={<Onboarding />} />
                    <Route path="/adaptation" element={<Adapting />} />
                    <Route path="/login" element={<Login />} />
                </Routes>
                <Footer />
            </div>
        </Router>
    );
}

export default App;
