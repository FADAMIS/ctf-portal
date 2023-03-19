<script>
    import { TeamFill, FlagFill, SettingsFill, UploadCloud2Fill, DownloadCloud2Fill, Message3Fill } from 'svelte-remixicon';
    let pages = [false, false, false, false, false, false];
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
        for (let i = 0; i < pages.length; i++) {
            pages[i] = false;
        }
        pages[4] = true;
    }
    function Export() {
        for (let i = 0; i < pages.length; i++) {
            pages[i] = false;
        }
        pages[5] = true;
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
                <div on:click={Import} class="flex w-52 transition-all hover:bg-violet-900 p-3 rounded-xl ml-2 mt-1">
                    <UploadCloud2Fill size="28px" class="fill-gray-300 mr-4"/>
                    <h1 class="text-xl text-gray-300 font-mono">Import</h1>
                </div>
                <div on:click={Export} class="flex w-52 transition-all hover:bg-violet-900 p-3 rounded-xl ml-2 mt-1">
                    <DownloadCloud2Fill size="28px" class="fill-gray-300 mr-4"/>
                    <h1 class="text-xl text-gray-300 font-mono">Export</h1>
                </div>
            </div>
            {#if pages[0] == true}
            <div class="w-full flex place-content-center">
                <h1 class="text-gray-300">General</h1>
            </div>
            {:else if pages[1] == true}
            <div class="w-full flex place-content-center">
                <h1 class="text-gray-300">Announcment</h1>
            </div>
            {:else if pages[2] == true}
            <div class="w-full flex place-content-center">
                <h1 class="text-gray-300">Challanges</h1>
            </div>
            {:else if pages[3] == true}
            <div class="w-full flex place-content-center">
                <h1 class="text-gray-300">Teams</h1>
            </div>
            {:else if pages[4] == true}
            <div class="w-full flex place-content-center">
                <h1 class="text-gray-300">Import</h1>
            </div>
            {:else if pages[5] == true}
            <div class="w-full flex place-content-center">
                <h1 class="text-gray-300">Export</h1>
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