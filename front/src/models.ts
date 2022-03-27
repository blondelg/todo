export class Task {
    id: number;
    title: string;
    done: boolean;

    deserialize(input: object) {
        Object.assign(this, input);
        return this;
    }
}

export class Toast {
    title: string;
    subtitle: string;
    body: string;

    constructor () {
        this.title = '';
        this.subtitle = '';
        this.body = '';
    }

    setTitle(title: string) {
        this.title = title;
    }

    setSubTitle(subtitle: string) {
        this.subtitle = subtitle;
    }

    setBody(body: string) {
        this.body = body;
    }
}