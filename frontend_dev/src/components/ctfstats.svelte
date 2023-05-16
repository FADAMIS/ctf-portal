<script>

    async function validate() {
          const response = await fetch('/api/points');
          if (response.status === 401) {
            window.location.href = '/';
          }
    }

    validate();

    import 'https://cdn.jsdelivr.net/npm/dohjs@latest/dist/doh.min.js';
    import { dataset_dev } from 'svelte/internal';

    let days = "00";
    let hours = "00";
    let minutes = "00";
    let seconds = "00";
    let nano = "00";

    function getTimer() {
        return fetch('/api/timedstart').then(res => res.json()).then(data => {   
            countDown(data)
        });
    }

    getTimer()

    var end = 0;

    function countDown(data) {
        if (end == 0) {
            end = data.stoptime;
        }
        let now = new Date().getTime().toString().slice(0, 10);
        let distance = end - parseInt(now);

        if (distance > 0) {
            days = Math.floor(distance / (24 * 60 * 60));
            hours = Math.floor((distance % (24 * 60 * 60)) / (60 * 60));
            minutes = Math.floor((distance % (60 * 60)) / (60));
            seconds = Math.floor((distance % (60)));
        }

        else if (distance < 0) {
            days = "00";
            hours = "00";
            minutes = "00";
            seconds = "00";
            nano = "00";
        }
        
        setTimeout(countDown, 100);
    }

    let announcements = [];

    fetch('/api/announcement').then(res => res.json())
        .then(data => {
            announcements = data.announcements})

    let teams = [];

    fetch('/api/points').then(res => res.json())
            .then(data => { teams = data.allpoints })

</script>

<main>
<div class="overflow-hidden">
    <div class="w-64 h-16 pt-5 pl-10 pr-5 hover:h-52 hover:w-96 transition-all border border-violet-500 rounded-xl bg-gray-950 overflow-y-scroll overflow-x-hidden top-[-10px] left-[-10px] z-10 absolute">
    <h1 class="text-white text-2xl mb-2">Announcements:</h1>
    {#each announcements as announcement}
    <h1 class="text-white break-word font-mono">admin: {announcement.message}</h1>
    {/each}
    </div>
    <div class="w-64 h-16 pt-5 pr-10 pl-5 hover:h-52 hover:w-96 transition-all border border-violet-500 rounded-xl bg-gray-950 overflow-y-scroll overflow-x-hidden top-[-10px] right-[-10px] z-10 absolute">
    <h1 class="text-white text-2xl mb-2">Leaderboard:</h1>
    {#each teams as team, order}
    <h1 class="text-white break-word font-mono">{order + 1}#_{team.team}</h1>
    {/each}
    </div>
    <div class="grid place-items-center h-20 w-screen bottom-3 z-10 absolute">
        <div class="text-gray-500 font-mono mb-1 text-sm">TIME LEFT</div>
        <div class="text-gray-100 font-mono text-4xl">{days}:{hours}:{minutes}:{seconds}</div>
        <div class="text-gray-500 font-mono mb-1 text-sm flex">
        <p class="text-xs ml-8 mr-8">days</p>
        <p class="text-xs mr-6">hours</p>
        <p class="text-xs mr-4">minutes</p>
        <p class="text-xs mr-4">seconds</p>

        </div>
    </div>
</div>
</main>