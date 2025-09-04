export namespace main {
	
	export class Character {
	    id: number;
	    name: string;
	    clan: string;
	    generation: string;
	    concept: string;
	    sire: string;
	    attributes: Record<string, number>;
	    // Go type: struct { Talents map[string]int "json:\"talents\""; Skills map[string]int "json:\"skills\""; Knowledges map[string]int "json:\"knowledges\"" }
	    abilities: any;
	    disciplines: Record<string, number>;
	    notes: string;
	
	    static createFrom(source: any = {}) {
	        return new Character(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.clan = source["clan"];
	        this.generation = source["generation"];
	        this.concept = source["concept"];
	        this.sire = source["sire"];
	        this.attributes = source["attributes"];
	        this.abilities = this.convertValues(source["abilities"], Object);
	        this.disciplines = source["disciplines"];
	        this.notes = source["notes"];
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
	export class Combatant {
	    id: number;
	    name: string;
	    initiativeBonus: number;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new Combatant(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.initiativeBonus = source["initiativeBonus"];
	        this.type = source["type"];
	    }
	}
	export class NPC {
	    id: number;
	    name: string;
	    clan: string;
	    generation: string;
	    notes: string;
	    attributes: Record<string, number>;
	
	    static createFrom(source: any = {}) {
	        return new NPC(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.clan = source["clan"];
	        this.generation = source["generation"];
	        this.notes = source["notes"];
	        this.attributes = source["attributes"];
	    }
	}
	export class SessionLog {
	    id: number;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new SessionLog(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.content = source["content"];
	    }
	}

}

