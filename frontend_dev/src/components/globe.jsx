import React, { useState, useEffect, createElement } from 'react';
import Globe from 'react-globe.gl';

export default function InterfaceGlobe() {
    const [isHover, setIsHover] = useState(false);
    const [geoData, setGeoData] = useState([]);
    const [hoverD, setHoverD] = useState(null);

    const handleMouseEnter = () => {
        setIsHover(true);
     };
  
     const handleMouseLeave = () => {
        setIsHover(false);
     };

    useEffect(() => {
        fetch('./map.geojson').then(res => res.json()).then(data => setGeoData(data.features));
    }, []);

    let challenges = [];

    function fetchChallenges() {
        fetch('/api/challenges').then(res => res.json()).then(data => {
            if (data.message === 'CTF is not yet started') {
                document.getElementById('status').innerHTML = 'CTF is not yet started';
            }
            else {
                document.getElementById('status').innerHTML = 'CTF is running';
            }
            });
    }

    function isCountryUsed(country) {
        fetchChallenges();
        for (let i = 0; i < challenges.length; i++) {
            if (country === challenges[i].country.toUpperCase() === country) {
                return true;
            }
        }
        return false;
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
            <h1 id="status" style={{ position: 'absolute', top: 0, left: 0, right: 0, bottom: 0, margin: 'auto', backgroundColor: 'rgb(255, 255, 255)', color: 'whitesmoke' }}>test</h1>
            <div id="challengeBlock" style={{ position: 'absolute', top: 0, left: 0, right: 0, bottom: 0, margin: 'auto', backgroundColor: 'rgb(18, 18, 18)', display: 'none', width: 600, height: 400, fontFamily: 'monospace', borderRadius: '10px', padding: '30px', border: '2px solid #5D34B3' }}>
                <div style={{ display: 'flex' }}>
                    <div>
                        <h1 id="challengeName" className="challengeInfo" style={{ color: 'whitesmoke', fontSize: '30px' }}>Challenge</h1>
                        <h3 id="challengeCountry" className="challengeInfo" style={{ color: 'whitesmoke', marginBottom: '20px' }}></h3>
                        <p id="challengeDescription" className="challengeInfo" style={{ color: 'gray' }}></p>
                    </div>
                    <div style={{ display: 'grid', gap: '10px', marginLeft: '40px', marginTop: '50px' }}>
                        <input id="challengeFlag" type="text" className="challengeInfo" placeholder='Flag here' style={{ width: '14rem', height: '2.2rem', color: '#5D34B3', padding: '5px', backgroundColor: 'rgb(16, 16, 16)', borderRadius: '10px', border: '2px solid #5D34B3' }}/>
                        <button id="challengeSubmit" className="challengeInfo" style={{ width: '14rem', height: '2.2rem', backgroundColor: '#5D34B3', borderRadius: '10px' }}>Submit</button>
                        <button id="challengeFiles" className="challengeInfo" style={{ width: '14rem', height: '2.2rem', backgroundColor: '#5D34B3', borderRadius: '10px' }}>Download Files</button>
                    </div>
                </div>
            </div>
      </div>   
    );
}
