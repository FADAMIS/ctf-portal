<script>
    let pages = [false, false, false, false];
    let teams = [{name: '', password: '', score: 0, solved: []}];
    let challanges = [];
    let files = [];
    let makeChallange = [{name: '', files: files, flag: '', points: 0, description: '', country: ''}];
    let challangeIndex = 0;
    let fileIndex = 0;
    let fileName;
    let teamName = '';
    let teamPassword = '';
    let timeStart;
    let timeEnd;
    let dateStart;
    let dateEnd;
    let announcement;
    let announcementId = 0;
    let teamIndex = 0;
    function General() {
        for (let i = 0; i < pages.length; i++) {
            pages[i] = false;
        }
        pages[0] = true;
    }
    function Announcment() {
        for (let i = 0; i < pages.length; i++) {
            pages[i] = false;
        }
        pages[1] = true;
    }
    function Challanges() {
        for (let i = 0; i < pages.length; i++) {
            pages[i] = false;
        }
        pages[2] = true;
    }
    function Teams() {
        for (let i = 0; i < pages.length; i++) {
            pages[i] = false;
        }
        pages[3] = true;
        fetch('/teams')
            .then(res => res.json())
            .then(data => {
                teams = data.users;
            })
    }
    function Import() {
        document.getElementById('import').click();
    }

    function addTime() {
        let start = new Date(dateStart + ' ' + timeStart).getTime() / 1000;
        let end = new Date(dateEnd + ' ' + timeEnd).getTime() / 1000;
        if (start < end) {
            fetch('/timedstart', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    starttime: start,
                    stoptime: end
                })
            }).then(res => {
                if (res.status == 200) {
                    alert('Time added')
                } else {
                    alert('💀')
                }
            })
        }
        else if (start > end) {
            alert('Start time must be before end time')
        }
        else if (start == end) {
            alert('CTF has to run at leats 1 second')
        }
        else {
            alert('Enter a valid time')
        }
    }

    function manualStart() {
        fetch('/manualstart', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                isstarted: true
            })
        }).then(res => {
            if (res.status == 200) {
                alert('CTF started')
            } else {
                alert('💀')
            }
        })
    }

    function manualStop() {
        fetch('/manualstart', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                isstarted: false
            })
        }).then(res => {
            if (res.status == 200) {
                alert('CTF stopped')
            } else {
                alert('💀')
            }
        })
    }

    function postAnnouncment() {
        fetch('/announcement')
            .then(res => res.json())
            .then(data => {
                announcementId = data.announcements.length})
            if (announcementId == 0) {
                announcementId = 0
            }
            else {
                announcementId = announcementId + 1
            }
        fetch('/announcement', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                message: announcement,
                id: announcementId
            })
        }).then(res => {
            if (res.status == 200) {
                alert('Announcment posted')
            } else {
                alert('💀')
            }
        })
    }

    function challangeFileUpload() {
        document.getElementById('makeChallangeFile').click();
        const fileInput = document.querySelector('#makeChallangeFile');

        fileInput.addEventListener('change', (e) => {
            const file = e.target.files[0];

            const reader = new FileReader();
            reader.onloadend = () => {
            const base64String = reader.result.replace('data:', '').replace(/^.+,/, '');

        files.push({filename: file.name, base64: base64String});
        console.log(makeChallange[0])
    };
    reader.readAsDataURL(file);
});

        fileIndex += 1;
        files = files;
    }

    function addChallange() {
        makeChallange = makeChallange;
        fetch('/upload', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: makeChallange[0].name,
                files: makeChallange[0].files,
                flag: makeChallange[0].flag,
                points: makeChallange[0].points,
                description: makeChallange[0].description,
                country: makeChallange[0].country
            })
        }).then(res => {
            if (res.status == 200) {
                alert('Challange added')
            } else {
                alert('💀')
            }
        })
    }

    function deleteChallange(challangeIndex) {
        fetch('/upload', {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: challanges[challangeIndex].name
            })
        }).then(res => {
            if (res.status == 200) {
                alert('Challange deleted')
            } else {
                alert('💀')
            }
        }).then(res =>{
            fetch('/challenges')
                .then(res => res.json())
                .then(data => {
                    challanges = data.challanges;
            })
        })
    }

    function addTeam() {
        fetch('/register', {

            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username: teamName,
                password: teamPassword
            })
        }).then(res => {
            if (res.status == 200) {
                alert('Team added')
            } else {
                alert('Team already exists')
            }
        }).then(res =>{
            fetch('/teams')
                .then(res => res.json())
                .then(data => {
                    teams = data.users;
            })
        })
    }

    function deleteTeam(teamIndex) {
        var delUsername = teams[teamIndex].username
        fetch('/teams', {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username: delUsername
            })
        }).then(res => {
            if (res.status == 200) {
                alert('Team deleted')
            } else {
                alert('💀')
            }
        }).then(res => {
            fetch('/teams')
                .then(res => res.json())
                .then(data => {
                    teams = data.users;
            })
        })
    }
</script>
<main>
    <div id="body" class="w-full h-full overflow-hidden">
        <div id="header" class="grid place-content-center w-screen h-20 bg-black outline outline-violet-500">
            <h1 class="text-center text-4xl font-mono text-gray-300">Admin page</h1>
        </div>
        <div class="flex">
            <div class="grid place-content-start w-56 h-screen bg-black">
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <div on:click={General} class="flex w-52 transition-all hover:bg-violet-900 p-3 rounded-xl ml-2 mt-10">
                    <h1 class="text-xl text-gray-300 font-mono">General</h1>
                </div>
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <div on:click={Announcment} class="flex w-52 transition-all hover:bg-violet-900 p-3 rounded-xl ml-2 mt-1">
                    <h1 class="text-xl text-gray-300 font-mono">Announcment</h1>
                </div>
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <div on:click={Challanges} class="flex w-52 transition-all hover:bg-violet-900 p-3 rounded-xl ml-2 mt-1">
                    <h1 class="text-xl text-gray-300 font-mono">Challanges</h1>
                </div>
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <div on:click={Teams} class="flex w-52 transition-all hover:bg-violet-900 p-3 rounded-xl ml-2 mt-1">
                    <h1 class="text-xl text-gray-300 font-mono">Teams</h1>
                </div>
            </div>
            {#if pages[0] == false && pages[1] == false && pages[2] == false && pages[3] == false}
            <div class="text-white">Later...</div>
            {:else if pages[0] == true} 
            <div class="w-full h-full flex place-content-center flex-col">
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">General</h1>
                <div class="w-full flex place-content-center mt-10">
                    <div class="grid grid-cols-1">
                        <label for="start" class="text-gray-300 text-xl font-mono ml-5">Start time</label>
                        <input type="date" class="w-40 h-10 ml-5 rounded-xl outline-none bg-gray-800 text-gray-300 flex place-content-center" bind:value={dateStart}>
                        <input type="time" class="w-40 h-10 ml-5 rounded-xl outline-none bg-gray-800 text-gray-300 flex place-content-center mt-3" bind:value={timeStart}>
                    </div>
                    <div class="grid grid-cols-1">
                        <label for="end" class="text-gray-300 text-xl font-mono ml-5">End time</label>
                        <input type="date" class="w-40 h-10 ml-5 rounded-xl outline-none bg-gray-800 text-gray-300 flex place-content-center" bind:value={dateEnd}>
                        <input type="time" class="w-40 h-10 ml-5 rounded-xl outline-none bg-gray-800 text-gray-300 flex place-content-center mt-3" bind:value={timeEnd}> 
                    </div>
                </div>
                <div class="w-full flex place-content-center">
                    <button class="w-40 h-10 mt-10 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono" on:click={addTime}>Save</button>
                </div>
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">Manual control:</h1>
                <div class="w-full flex place-content-center mt-4">
                    <button class="w-40 h-10 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono" on:click={manualStop}>Stop</button>
                    <button class="w-40 h-10 ml-5 rounded-xl bg-violet-800 hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono" on:click={manualStart}>Start</button>
                </div>
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">Import/Export Challanges:</h1>
                <div class="w-full h-full flex justify-center">
                    <input id='import' type='file' hidden />
                    <button class="w-40 h-10 mt-5 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono" on:click={Import}>Import</button>
                    <button class="w-40 h-10 mt-5 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono">Export</button>
                </div>
            </div>
            {:else if pages[1] == true}
            <div class="w-full h-full flex place-content-center flex-col">
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">Announcment</h1>
                <div class="w-full flex place-content-center mt-10">
                    <label for="end" class="text-gray-300 text-xl font-mono mt-1">Message: </label>
                    <input type="text" class="p-2 w-7/12 h-10 ml-5 rounded-xl border border-violet-500 bg-gray-800 text-gray-300" bind:value={announcement}>
                </div>
                <div class="w-full flex place-content-center">
                    <button class="w-40 h-10 mt-5 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono" on:click={postAnnouncment}>Send</button>
                </div>
            </div>
            {:else if pages[2] == true}
            <div class="w-full h-full flex place-content-center flex-col">
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">Challanges</h1>
                <div class="w-full p-2 lg:pl-64 lg:pr-64">
                    {#each challanges as challange}
                    <div class="w-full bg-black border border-violet-500 rounded-xl mt-2 p-10 pt-7">
                        <h1 class="text-white text-2xl font-mono">Name:</h1>
                        <input type="text" class="w-full h-10 rounded-xl p-2 bg-gray-800 text-gray-50 flex place-content-center" value={challange.name} disabled>
                        <div class="flex">
                            <div>
                                <h1 class="text-white text-2xl font-mono">Flag:</h1>
                                <input type="text" class="w-full h-10 rounded-xl p-2 bg-gray-800 text-gray-50 flex place-content-center" value={challange.description} disabled>
                            </div>
                            <div class="ml-2">
                                <h1 class="text-white text-2xl font-mono">Points:</h1>
                                <input type="text" class="w-48 h-10 rounded-xl p-2 bg-gray-800 text-gray-50 flex place-content-center text-center" value={challange.points} disabled>
                            </div>
                        </div>
                        <h1 class="text-white text-xl font-mono mt-3">Files:</h1>
                        {#each challange.files as file}
                        <p class="text-white">{file.name}</p>
                        {/each}
                        <div class="w-full flex place-content-end">
                            <button class="w-40 h-10 mt-5 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono">Delete</button>
                        </div>
                    </div>
                    {/each}
                    <div class="flex justify-center mt-10">
                    </div>
                    <div class="w-3/4 ml-10 mt-8">
                        <div class="flex">
                            <div class="w-40">
                                <h1 class="text-white text-2xl font-mono">Name:</h1>
                                <input type="text" class="w-full h-10 rounded-xl p-2 bg-gray-800 text-gray-50 flex place-content-center" bind:value={makeChallange[0].name}>
                            </div>
                            <div class="ml-2">
                                <h1 class="text-white text-2xl font-mono">Description:</h1>
                                <input type="text" class="w-full h-10 rounded-xl p-2 bg-gray-800 text-gray-50 flex place-content-center" bind:value={makeChallange[0].description}>
                            </div>
                            <div class="ml-2">
                                <h1 class="text-white text-2xl font-mono">Country:</h1>
                                <input type="text" class="w-full h-10 rounded-xl p-2 bg-gray-800 text-gray-50 flex place-content-center" bind:value={makeChallange[0].country}>
                            </div>
                        </div>
                        <div class="flex">
                            <div>
                                <h1 class="text-white text-2xl font-mono">Flag:</h1>
                                <input type="text" class="w-full h-10 rounded-xl p-2 bg-gray-800 text-gray-50 flex place-content-center" bind:value={makeChallange[0].flag}>
                            </div>
                            <div class="ml-2">
                                <h1 class="text-white text-2xl font-mono">Points:</h1>
                                <input type="number" class="w-48 h-10 rounded-xl p-2 bg-gray-800 text-gray-50 flex place-content-center text-center" bind:value={makeChallange[0].points}>
                            </div>
                        </div>
                        <h1 class="text-white text-xl font-mono mt-3">Files:</h1>
                        {#each files as file}
                        <p class="text-white">{file.name}</p>
                        {/each}
                        <div class="w-full flex place-content-end">
                            <input type="file" id="makeChallangeFile" hidden>
                            <button class="w-40 h-10 mt-5 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono" on:click={challangeFileUpload}>Upload Files</button>
                            <button class="w-40 h-10 mt-5 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono" on:click={addChallange}>Save</button>
                        </div>
                    </div>
                </div>
            </div>
            {:else if pages[3] == true}
            <div class="w-full h-full flex place-content-center flex-col">
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">Teams</h1>
                <div class="w-full p-2 lg:pl-64 lg:pr-64">
                            {#each teams as team}
                            <div class="w-full h-64 bg-black border border-violet-500 rounded-xl mt-2 p-10 pt-7">
                                <h1 class="text-white text-2xl font-mono">Username:</h1>
                                <input type="text" class="w-full h-10 rounded-xl p-2 bg-gray-800 text-gray-50 flex place-content-center" value={team.username} disabled>
                                <p class="text-white text-2xl font-mono">Password:</p>
                                <input type="text" class="w-full h-10 rounded-xl p-2 bg-gray-800 text-gray-50 flex place-content-center" value={team.password} disabled>
                                <div class="w-full flex place-content-end">
                                    <button class="w-40 h-10 mt-5 ml-5 rounded-xl bg-violet-700 hover:bg-violet-900 transition-all text-gray-300 flex place-content-center p-2 font-mono" on:click={() => deleteTeam(teamIndex)}>Delete</button>
                                    {teamIndex += 1}
                                </div>
                            </div>
                            {/each}
                            <div class="flex justify-center mt-10">
                            </div>
                    <div class="w-full flex place-content-center mt-10">
                        <div class="grid grid-cols-1">
                            <label for="start" class="text-gray-300 text-xl font-mono ml-5">Team name</label>
                            <input type="username" class="w-40 h-10 ml-5 rounded-xl p-2 bg-gray-800 text-gray-300 flex place-content-center" bind:value={teamName}>
                        </div>
                        <div class="grid grid-cols-1">
                            <label for="end" class="text-gray-300 text-xl font-mono ml-5">Team password</label>
                            <input type="password" class="w-40 h-10 ml-5 rounded-xl p-2 bg-gray-800 text-gray-300 flex place-content-center" bind:value={teamPassword}>
                        </div>
                    </div>
                    <div class="w-full flex place-content-center">
                        <button class="w-40 h-10 mt-10 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono" on:click={addTeam}>Add</button>
                    </div>
                </div>
            </div>
            {/if}
        </div>
    </div>
</main>
<style>
    #body {
        background-color: #161616;
        opacity: 0.8;
        background-size: 10px 10px;
        background-image: repeating-linear-gradient(45deg, #2e2e2e 0, #2e2e2e 1px, #040404 0, #010101 50%);
    }
</style>