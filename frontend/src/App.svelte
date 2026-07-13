<script lang="ts">
    import { writable, derived } from 'svelte/store';
    import { SelectClock, SelectCountdown } from '../wailsjs/go/main/App.js';
    import { main as m } from '../wailsjs/go/models';
    import { EventsOn } from '../wailsjs/runtime';

    const debug = true;
    const data = writable<m.Data>(new m.Data());

    EventsOn('data', (incoming: m.Data) => {
        if (debug) {
            console.log(JSON.stringify(incoming, null, 4));
        }
        data.set(incoming);
    });

    function selectClock(selected: m.Clock): void {
        if (selected.selected) {
            SelectClock('');
            selected.selected = false;
        } else {
            SelectClock(selected.tz);
            selected.selected = true;
        }

        data.update((cur) => {
            return cur;
        });
    }

    function selectCountdown(selected: m.Countdown): void {
        if (selected.selected) {
            selected.selected = false;
            SelectCountdown(0);
        } else {
            selected.selected = true;
            SelectCountdown(selected.value);
        }

        data.update((cur) => {
            return cur;
        });
    }
</script>

<main>
    <section class="endsIn">
        Class ends {$data.endsin}
    </section>

    <section class="clockSection">
        {#if $data.clocks}
            {#each $data.clocks as clock}
                <button class="btn" on:click={() => selectClock(clock)}>
                    <article class="clock" class:selected={clock.selected}>
                        <div>{clock.tz}</div>
                        <div>{clock.value}</div>
                    </article>
                </button>
            {/each}
        {/if}
    </section>

    <section class="countdownSection">
        {#if $data.countdowns}
            {#each $data.countdowns as countdown}
                <button class="btn" on:click={() => selectCountdown(countdown)}>
                    <article class="countdown" class:selected={countdown.selected}>
                        {countdown.value} mins
                    </article>
                </button>
            {/each}
        {/if}
    </section>

    <section class="countDown">
        {#if $data.segmented}
            <span>Back in:</span>

            <section>
                {`${$data.segmented.hours}`}
                : {$data.segmented.mins}
                : {$data.segmented.secs}
            </section>
        {/if}
    </section>
</main>

<style>
    main {
        padding-block-start: 1rem;
        display: flex;
        flex-direction: column;
        gap: 2rem;
    }

    section {
        font-weight: var(--font-weight-5);
        font-size: var(--font-size-4);
        padding-inline: 1rem;
    }

    .endsIn {
        color: grey;
        display: flex;
        justify-content: center;
        font-size: small;
    }

    .countDown {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-grow: 1;
        gap: 2rem;

        & section {
            display: flex;
            justify-content: center;
            font-size: var(--size-relative-10);
            border-width: var(--border-size-1);
            border-color: var(--green-1);
            box-shadow: var(--shadow-2);
            color: var(--green-8);
            width: var(--size-sm);
        }
    }

    .clockSection {
        display: grid;
        gap: 1rem;

        & .clock {
            display: grid;
            grid-template-columns: 3fr 5fr;
        }
    }

    .countdownSection {
        display: flex;
        justify-content: center;
        align-items: flex-start;

        & .countdown {
            font-size: var(--font-size-1);
        }
    }

    article {
        box-shadow: var(--inner-shadow-1);
        padding: 1rem;

        &:hover {
            background-color: var(--green-0);
            color: var(--green-6);
            box-shadow: var(--inner-shadow-3);
        }
    }

    .selected {
        background-color: var(--green-1);
        color: var(--green-8);
        box-shadow: var(--inner-shadow-3);
        font-size: var(--font-size-5);
    }

    .btn {
        background-color: transparent;
    }
</style>
