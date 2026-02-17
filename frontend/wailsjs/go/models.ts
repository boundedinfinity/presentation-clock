export namespace main {
	
	export class Clock {
	    tz: string;
	    value: string;
	    selected: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Clock(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tz = source["tz"];
	        this.value = source["value"];
	        this.selected = source["selected"];
	    }
	}
	export class Countdown {
	    value: number;
	    selected: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Countdown(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.value = source["value"];
	        this.selected = source["selected"];
	    }
	}
	export class SegmentedDuration {
	    hour: number;
	    hours: string;
	    min: number;
	    mins: string;
	    sec: number;
	    secs: string;
	
	    static createFrom(source: any = {}) {
	        return new SegmentedDuration(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hour = source["hour"];
	        this.hours = source["hours"];
	        this.min = source["min"];
	        this.mins = source["mins"];
	        this.sec = source["sec"];
	        this.secs = source["secs"];
	    }
	}
	export class Data {
	    clocks: Clock[];
	    clock?: Clock;
	    countdowns: Countdown[];
	    countdown?: Countdown;
	    segmented?: SegmentedDuration;
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.clocks = this.convertValues(source["clocks"], Clock);
	        this.clock = this.convertValues(source["clock"], Clock);
	        this.countdowns = this.convertValues(source["countdowns"], Countdown);
	        this.countdown = this.convertValues(source["countdown"], Countdown);
	        this.segmented = this.convertValues(source["segmented"], SegmentedDuration);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

