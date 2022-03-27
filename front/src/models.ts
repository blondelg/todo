export class Task {
    id: number;
    title: string;
    done: boolean;

    deserialize(input: object) {
        Object.assign(this, input);
        return this;
    }
}