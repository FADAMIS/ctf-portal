
import React, { useState, useEffect, createElement } from 'react';
import Globe from 'react-globe.gl';

let challenges = [];

function isCountryUsed(country) {
        for (let i = 0; i < challenges.length; i++) {
            if (country === challenges[i].country.toUpperCase()) {
                return true;
            }
        }
        return false;
    }

function fetchChallenges() {
        fetch('/api/challenges').then(res => res.json()).then(data => {
            if (data.status === 401) {
                alert('CTF is not yet started');
            }
            else {
                challenges = data.challenges;
            }
        });
    }

const types = {
        //   File Extension   MIME Type
          'abs':           'audio/x-mpeg',
          'ai':            'application/postscript',
          'aif':           'audio/x-aiff',
          'aifc':          'audio/x-aiff',
          'aiff':          'audio/x-aiff',
          'aim':           'application/x-aim',
          'art':           'image/x-jg',
          'asf':           'video/x-ms-asf',
          'asx':           'video/x-ms-asf',
          'au':            'audio/basic',
          'avi':           'video/x-msvideo',
          'avx':           'video/x-rad-screenplay',
          'bcpio':         'application/x-bcpio',
          'bin':           'application/octet-stream',
          'bmp':           'image/bmp',
          'body':          'text/html',
          'cdf':           'application/x-cdf',
          'cer':           'application/pkix-cert',
          'class':         'application/java',
          'cpio':          'application/x-cpio',
          'csh':           'application/x-csh',
          'css':           'text/css',
          'dib':           'image/bmp',
          'doc':           'application/msword',
          'dtd':           'application/xml-dtd',
          'dv':            'video/x-dv',
          'dvi':           'application/x-dvi',
          'eot':           'application/vnd.ms-fontobject',
          'eps':           'application/postscript',
          'etx':           'text/x-setext',
          'exe':           'application/octet-stream',
          'gif':           'image/gif',
          'gtar':          'application/x-gtar',
          'gz':            'application/x-gzip',
          'hdf':           'application/x-hdf',
          'hqx':           'application/mac-binhex40',
          'htc':           'text/x-component',
          'htm':           'text/html',
          'html':          'text/html',
          'ief':           'image/ief',
          'jad':           'text/vnd.sun.j2me.app-descriptor',
          'jar':           'application/java-archive',
          'java':          'text/x-java-source',
          'jnlp':          'application/x-java-jnlp-file',
          'jpe':           'image/jpeg',
          'jpeg':          'image/jpeg',
          'jpg':           'image/jpeg',
          'js':            'application/javascript',
          'jsf':           'text/plain',
          'json':          'application/json',
          'jspf':          'text/plain',
          'kar':           'audio/midi',
          'latex':         'application/x-latex',
          'm3u':           'audio/x-mpegurl',
          'mac':           'image/x-macpaint',
          'man':           'text/troff',
          'mathml':        'application/mathml+xml',
          'me':            'text/troff',
          'mid':           'audio/midi',
          'midi':          'audio/midi',
          'mif':           'application/x-mif',
          'mov':           'video/quicktime',
          'movie':         'video/x-sgi-movie',
          'mp1':           'audio/mpeg',
          'mp2':           'audio/mpeg',
          'mp3':           'audio/mpeg',
          'mp4':           'video/mp4',
          'mpa':           'audio/mpeg',
          'mpe':           'video/mpeg',
          'mpeg':          'video/mpeg',
          'mpega':         'audio/x-mpeg',
          'mpg':           'video/mpeg',
          'mpv2':          'video/mpeg2',
          'ms':            'application/x-wais-source',
          'nc':            'application/x-netcdf',
          'oda':           'application/oda',
          'odb':           'application/vnd.oasis.opendocument.database',
          'odc':           'application/vnd.oasis.opendocument.chart',
          'odf':           'application/vnd.oasis.opendocument.formula',
          'odg':           'application/vnd.oasis.opendocument.graphics',
          'odi':           'application/vnd.oasis.opendocument.image',
          'odm':           'application/vnd.oasis.opendocument.text-master',
          'odp':           'application/vnd.oasis.opendocument.presentation',
          'ods':           'application/vnd.oasis.opendocument.spreadsheet',
          'odt':           'application/vnd.oasis.opendocument.text',
          'otg':           'application/vnd.oasis.opendocument.graphics-template',
          'oth':           'application/vnd.oasis.opendocument.text-web',
          'otp':           'application/vnd.oasis.opendocument.presentation-template',
          'ots':           'application/vnd.oasis.opendocument.spreadsheet-template',
          'ott':           'application/vnd.oasis.opendocument.text-template',
          'ogx':           'application/ogg',
          'ogv':           'video/ogg',
          'oga':           'audio/ogg',
          'ogg':           'audio/ogg',
          'otf':           'application/x-font-opentype',
          'spx':           'audio/ogg',
          'flac':          'audio/flac',
          'anx':           'application/annodex',
          'axa':           'audio/annodex',
          'axv':           'video/annodex',
          'xspf':          'application/xspf+xml',
          'pbm':           'image/x-portable-bitmap',
          'pct':           'image/pict',
          'pdf':           'application/pdf',
          'pgm':           'image/x-portable-graymap',
          'pic':           'image/pict',
          'pict':          'image/pict',
          'pls':           'audio/x-scpls',
          'png':           'image/png',
          'pnm':           'image/x-portable-anymap',
          'pnt':           'image/x-macpaint',
          'ppm':           'image/x-portable-pixmap',
          'ppt':           'application/vnd.ms-powerpoint',
          'pps':           'application/vnd.ms-powerpoint',
          'ps':            'application/postscript',
          'psd':           'image/vnd.adobe.photoshop',
          'qt':            'video/quicktime',
          'qti':           'image/x-quicktime',
          'qtif':          'image/x-quicktime',
          'ras':           'image/x-cmu-raster',
          'rdf':           'application/rdf+xml',
          'rgb':           'image/x-rgb',
          'rm':            'application/vnd.rn-realmedia',
          'roff':          'text/troff',
          'rtf':           'application/rtf',
          'rtx':           'text/richtext',
          'sfnt':          'application/font-sfnt',
          'sh':            'application/x-sh',
          'shar':          'application/x-shar',
          'sit':           'application/x-stuffit',
          'snd':           'audio/basic',
          'src':           'application/x-wais-source',
          'sv4cpio':       'application/x-sv4cpio',
          'sv4crc':        'application/x-sv4crc',
          'svg':           'image/svg+xml',
          'svgz':          'image/svg+xml',
          'swf':           'application/x-shockwave-flash',
          't':             'text/troff',
          'tar':           'application/x-tar',
          'tcl':           'application/x-tcl',
          'tex':           'application/x-tex',
          'texi':          'application/x-texinfo',
          'texinfo':       'application/x-texinfo',
          'tif':           'image/tiff',
          'tiff':          'image/tiff',
          'tr':            'text/troff',
          'tsv':           'text/tab-separated-values',
          'ttf':           'application/x-font-ttf',
          'txt':           'text/plain',
          'ulw':           'audio/basic',
          'ustar':         'application/x-ustar',
          'vxml':          'application/voicexml+xml',
          'xbm':           'image/x-xbitmap',
          'xht':           'application/xhtml+xml',
          'xhtml':         'application/xhtml+xml',
          'xls':           'application/vnd.ms-excel',
          'xml':           'application/xml',
          'xpm':           'image/x-xpixmap',
          'xsl':           'application/xml',
          'xslt':          'application/xslt+xml',
          'xul':           'application/vnd.mozilla.xul+xml',
          'xwd':           'image/x-xwindowdump',
          'vsd':           'application/vnd.visio',
          'wav':           'audio/x-wav',
          'wbmp':          'image/vnd.wap.wbmp',
          'wml':           'text/vnd.wap.wml',
          'wmlc':          'application/vnd.wap.wmlc',
          'wmls':          'text/vnd.wap.wmlsc',
          'wmlscriptc':    'application/vnd.wap.wmlscriptc',
          'wmv':           'video/x-ms-wmv',
          'woff':          'application/font-woff',
          'woff2':         'application/font-woff2',
          'wrl':           'model/vrml',
          'wspolicy':      'application/wspolicy+xml',
          'z':             'application/x-compress',
          'zip':           'application/zip'
      };

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

    function flagSubmit() {
      let flag = document.getElementById('challengeFlag').value;
      let challange = document.getElementById('challengeName').innerHTML;
      fetch('/api/validate', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          challenge: challange,
          value: flag
        })
      })
        .then(response => response.status)
        .then(status => {
          if (status === 200) {
            alert("Correct Flag!");
          } else if (status === 403) {
            alert("Already anwsered!");
          } else if (status === 409) {
            alert("Wrong Flag!");
          } else {
            alert("Something went wrong!");
          }

        })
    }

    function download() {
      let files = []
      for (let i = 0; i < challenges.length; i++) {
        if (challenges[i].name === document.getElementById('challengeName').innerHTML) {
          for (let j = 0; j < challenges[i].files.length; j++) {
            files.push({ filename : challenges[i].files[j].filename, base64 : challenges[i].files[j].base64 });
          }
        }
      }

      if files == [] {
	  return
      }

      for (let i = 0; i < files.length; i++) {
          var downloadLink = document.createElement("a");
          var type = "";
          for (let j = 0; j < types.length; j++) {
            if (files[i].filename.includes(types[i].extension)) {
              type = types[i].type;
            }
          }
          downloadLink.href = "data:" + type + ";base64," + files[i].base64;
          downloadLink.download = files[i].filename;
          downloadLink.click();
      }
    }

    const handlePolygonHover = (polygon) => {
        setHoverD(polygon);  
    }

    function getChallengeByCountry(country) {
	for (let i = 0; i < challenges.length; i++) {
            if (country === challenges[i].country.toUpperCase()) {
                return challenges[i];
            }
        }
    }

    function handlePolygonClick(d) {
        if (isCountryUsed(d.properties.ISO_A2) === true) {
	    let challenge = getChallengeByCountry(d.properties.ISO_A2);
            document.getElementById('challengeBlock').style.display = 'block';
	    document.getElementById('challengeName').innerHTML = challenge.name;
            document.getElementById('challengeCountry').innerHTML = d.properties.ADMIN;
            document.getElementById('challengeDescription').innerHTML = challenge.description;
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
            polygonCapColor={d => isCountryUsed(d.properties.ISO_A2) ? '#3d2473' : 'rgb(22, 22, 22)'}
            polygonSideColor = {() => 'rgb(93, 52, 179)'}
            polygonStrokeColor = {() => '#5D34B3'}
            onPolygonHover={handlePolygonHover}
            polygonAltitude={d => (hoverD && d.properties.ADMIN === hoverD.properties.ADMIN ? 0.04 : 0.02)}
            polygonsTransitionDuration={150}
            onPolygonClick={d => handlePolygonClick(d)}
            />

            <div id="challengeBlock" style={{ position: 'absolute', top: 0, left: 0, right: 0, bottom: 0, margin: 'auto', backgroundColor: 'rgb(18, 18, 18)', display: 'none', width: 600, height: 400, fontFamily: 'monospace', borderRadius: '10px', padding: '30px', border: '2px solid #5D34B3' }}>
                <div style={{ display: 'flex' }}>
                    <div>
                        <h1 id="challengeName" className="challengeInfo" style={{ color: 'whitesmoke', fontSize: '30px' }}>Challenge</h1>
                        <h3 id="challengeCountry" className="challengeInfo" style={{ color: 'whitesmoke', marginBottom: '20px' }}></h3>
                        <p id="challengeDescription" className="challengeInfo" style={{ color: 'gray' }}></p>
                    </div>
                    <div style={{ display: 'grid', gap: '10px', marginLeft: '40px', marginTop: '50px' }}>
                        <input id="challengeFlag" type="text" className="challengeInfo" placeholder='Flag here' style={{ width: '14rem', height: '2.2rem', color: '#5D34B3', padding: '5px', backgroundColor: 'rgb(16, 16, 16)', borderRadius: '10px', border: '2px solid #5D34B3' }}/>
                        <button id="challengeSubmit" className="challengeInfo" onClick={flagSubmit} style={{ width: '14rem', height: '2.2rem', backgroundColor: '#5D34B3', borderRadius: '10px' }}>Submit</button>
                        <button id="challengeFiles" className="challengeInfo" onClick={download} style={{ width: '14rem', height: '2.2rem', backgroundColor: '#5D34B3', borderRadius: '10px' }}>Download Files</button>
                    </div>
                </div>
            </div>
      </div>   
    );
}

fetchChallenges();
setInterval(fetchChallenges, 2*60*1000)
