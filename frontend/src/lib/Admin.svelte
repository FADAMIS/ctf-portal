<script>
    import { TeamFill, FlagFill, SettingsFill, UploadCloud2Fill, DownloadCloud2Fill, Message3Fill } from 'svelte-remixicon';
    let pages = [false, false, false, false];
    let teams = [];
    function General() {
        for (let i = 0; i < pages.length; i++) {
            pages[i] = false;
        }
        pages[0] = true;
        console.log(pages)
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
    }
    function Import() {
        document.getElementById('import').click();
    }
</script>
<main>
    <div id="body" class="w-full h-full overflow-hidden">
        <div id="header" class="grid place-content-center w-screen h-20 bg-black outline outline-violet-500">
            <h1 class="text-center text-4xl font-mono text-gray-300">Admin page</h1>
        </div>
        <div class="flex">
            <div class="grid place-content-start w-56 h-screen bg-black">
                <div on:click={General} class="flex w-52 transition-all hover:bg-violet-900 p-3 rounded-xl ml-2 mt-10">
                    <SettingsFill size="28px" class="fill-gray-300 mr-4"/>
                    <h1 class="text-xl text-gray-300 font-mono">General</h1>
                </div>
                <div on:click={Announcment} class="flex w-52 transition-all hover:bg-violet-900 p-3 rounded-xl ml-2 mt-1">
                    <Message3Fill size="28px" class="fill-gray-300 mr-4"/>
                    <h1 class="text-xl text-gray-300 font-mono">Announcment</h1>
                </div>
                <div on:click={Challanges} class="flex w-52 transition-all hover:bg-violet-900 p-3 rounded-xl ml-2 mt-1">
                    <FlagFill size="28px" class="fill-gray-300 mr-4"/>
                    <h1 class="text-xl text-gray-300 font-mono">Challanges</h1>
                </div>
                <div on:click={Teams} class="flex w-52 transition-all hover:bg-violet-900 p-3 rounded-xl ml-2 mt-1">
                    <TeamFill size="28px" class="fill-gray-300 mr-4"/>
                    <h1 class="text-xl text-gray-300 font-mono">Teams</h1>
                </div>
            </div>
            {#if pages[0] == false && pages[1] == false && pages[2] == false && pages[3] == false}
            <div class="w-full h-full flex place-content-center flex-col">
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">General</h1>
                <div class="w-full flex place-content-center mt-10">
                    <div class="grid grid-cols-1">
                        <label for="start" class="text-gray-300 text-xl font-mono ml-5">Start time</label>
                        <input type="time" class="w-40 h-10 ml-5 rounded-xl outline-none bg-gray-800 text-gray-300 flex place-content-center">
                    </div>
                    <div class="grid grid-cols-1">
                        <label for="end" class="text-gray-300 text-xl font-mono ml-5">End time</label>
                        <input type="time" class="w-40 h-10 ml-5 rounded-xl outline-none bg-gray-800 text-gray-300 flex place-content-center">
                    </div>
                </div>
                <div class="w-full flex place-content-center">
                    <button class="w-40 h-10 mt-10 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono">Save</button>
                </div>
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">Manual control:</h1>
                <div class="w-full flex place-content-center mt-4">
                    <button class="w-40 h-10 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono">Stop</button>
                    <button class="w-40 h-10 ml-5 rounded-xl bg-violet-800 hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono">Start</button>
                </div>
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">Import/Export Challanges:</h1>
                <div class="w-full h-full flex justify-center">
                    <input id='import' type='file' hidden />
                    <button class="w-40 h-10 mt-5 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono" on:click={Import}>Import</button>
                    <button class="w-40 h-10 mt-5 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono">Export</button>
                </div>
            </div>
            {:else if pages[0] == true} 
            <div class="w-full h-full flex place-content-center flex-col">
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">General</h1>
                <div class="w-full flex place-content-center mt-10">
                    <div class="grid grid-cols-1">
                        <label for="start" class="text-gray-300 text-xl font-mono ml-5">Start time</label>
                        <input type="time" class="w-40 h-10 ml-5 rounded-xl outline-none bg-gray-800 text-gray-300 flex place-content-center">
                    </div>
                    <div class="grid grid-cols-1">
                        <label for="end" class="text-gray-300 text-xl font-mono ml-5">End time</label>
                        <input type="time" class="w-40 h-10 ml-5 rounded-xl outline-none bg-gray-800 text-gray-300 flex place-content-center">
                    </div>
                </div>
                <div class="w-full flex place-content-center">
                    <button class="w-40 h-10 mt-10 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono">Save</button>
                </div>
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">Manual control:</h1>
                <div class="w-full flex place-content-center mt-4">
                    <button class="w-40 h-10 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono">Stop</button>
                    <button class="w-40 h-10 ml-5 rounded-xl bg-violet-800 hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono">Start</button>
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
                    <input type="text" class="p-2 w-7/12 h-10 ml-5 rounded-xl border border-violet-500 bg-gray-800 text-gray-300">
                </div>
                <div class="w-full flex place-content-center">
                    <button class="w-40 h-10 mt-5 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono">Send</button>
                </div>
            </div>
            {:else if pages[2] == true}
            <div class="w-full h-full flex place-content-center flex-col">
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">Challanges</h1>
            </div>
            {:else if pages[3] == true}
            <div class="w-full h-full flex place-content-center flex-col">
                <h1 class="text-gray-300 text-2xl text-center mt-5 font-mono">Teams</h1>
                <div class="w-full p-2 lg:pl-64 lg:pr-64">
                            {#each teams as team}
                            <div class="w-full h-52 bg-black border border-violet-500 rounded-xl mt-2">
                                <p class="text-white">test</p>
                            </div>
                            {/each}
                    <div class="w-full flex place-content-center mt-10">
                        <div class="grid grid-cols-1">
                            <label for="start" class="text-gray-300 text-xl font-mono ml-5">Team name</label>
                            <input type="username" class="w-40 h-10 ml-5 rounded-xl p-2 bg-gray-800 text-gray-300 flex place-content-center">
                        </div>
                        <div class="grid grid-cols-1">
                            <label for="end" class="text-gray-300 text-xl font-mono ml-5">Team password</label>
                            <input type="password" class="w-40 h-10 ml-5 rounded-xl p-2 bg-gray-800 text-gray-300 flex place-content-center">
                        </div>
                    </div>
                    <div class="w-full flex place-content-center">
                        <button class="w-40 h-10 mt-10 ml-5 rounded-xl bg-black hover:bg-violet-700 transition-all text-gray-300 flex place-content-center p-2 font-mono">Add</button>
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