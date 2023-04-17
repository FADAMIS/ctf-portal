import React, { useState, useEffect, createElement } from 'react';
import Globe from 'react-globe.gl';

export default function InterfaceGlobe() {
    const [geoData, setGeoData] = useState([]);
    const [hoverD, setHoverD] = useState(null);

    useEffect(() => {
        fetch('./map.geojson').then(res => res.json()).then(data => setGeoData(data.features));
    }, []);

    function fetchChallenges() {
        fetch('http://localhost:5000/challenges', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            }.then(res => res.json()).then(data => {
                console.log(data);
            })
        })
    }

    function isCountryUsed(country) {
        if (country === 'United States of America') {
            return true;
        }
        else {
            return false;
        }
    }

    const handlePolygonHover = (polygon) => {
        setHoverD(polygon);  
    }

    function handlePolygonClick(d) {
        console.log(d.properties.ADMIN);
        if (isCountryUsed(d.properties.ADMIN) === true) {
            document.getElementById('challengeBlock').style.display = 'block';
            document.getElementById('challengeCountry').innerHTML = d.properties.ADMIN;
            document.getElementById('challengeDescription').innerHTML = 'This is a challenge description';
            addEventListener('click', () => {
                if(event.target.id !== 'challengeBlock' && event.target.className !== 'challengeInfo') {
                    document.getElementById('challengeBlock').style.display = 'none';
                }
            });
        }
        else {
            document.getElementById('challengeBlock').style.display = 'none';
        }
    }

    return (
        <div style={{position: 'absolute', zIndex: 1 }}>
            <Globe
            globeImageUrl = '//unpkg.com/three-globe/example/img/earth-dark.jpg'
            backgroundImageUrl = '//unpkg.com/three-globe/example/img/night-sky.png'
            polygonsData = {geoData.filter(d => d.properties.ISO_A2 !== 'AQ')}
            polygonCapColor={d => isCountryUsed(d.properties.ADMIN) ? '#3d2473' : 'rgb(22, 22, 22)'}
            polygonSideColor = {() => 'rgb(93, 52, 179)'}
            polygonStrokeColor = {() => '#5D34B3'}
            onPolygonHover={handlePolygonHover}
            polygonAltitude={d => (hoverD && d.properties.ADMIN === hoverD.properties.ADMIN ? 0.04 : 0.02)}
            polygonsTransitionDuration={150}
            onPolygonClick={d => handlePolygonClick(d)}
            />
            <div id="challengeBlock" style={{ position: 'absolute', top: 0, left: 0, right: 0, bottom: 0, margin: 'auto', backgroundColor: 'rgb(22, 22, 22)', display: 'none', width: 600, height: 400, fontFamily: 'monospace', borderRadius: '10px', padding: '30px', border: '2px solid #5D34B3' }}>
                <h1 id="challengeName" className="challengeInfo" style={{ color: 'whitesmoke' }}>Challenge</h1>
                <h3 id="challengeCountry" className="challengeInfo" style={{ color: 'whitesmoke' }}></h3>
                <p id="challengeDescription" className="challengeInfo" style={{ color: 'whitesmoke' }}></p>
            </div>
      </div>   
    );
}