// Concept 1: Generic React Hook
function useState<T>(initial: T): [T, (value: T) => void] {
    let state = initial;
    const setState = (newState: T) => { state = newState; };
    return [state, setState];
}

// Concept 2: Render Props Pattern
interface ListProps<T> {
    items: T[];
    renderItem: (item: T, index: number) => React.ReactNode;
}

// Concept 3: Higher-Order Component
function withLogging<P extends object>(Component: React.ComponentType<P>): React.ComponentType<P> {
    return (props: P) => {
        console.log(`Rendering ${Component.displayName}`);
        return <Component {...props} />;
    };
}

// Concept 4: Custom Hook Pattern
function useLocalStorage<T>(key: string, initialValue: T): [T, (value: T) => void] {
    const readValue = (): T => {
        const item = localStorage.getItem(key);
        return item ? JSON.parse(item) : initialValue;
    };
    const setValue = (value: T): void => {
        localStorage.setItem(key, JSON.stringify(value));
    };
    return [readValue(), setValue];
}

// Concept 5: Context with TypeScript
interface ThemeContextType {
    theme: "light" | "dark";
    toggleTheme: () => void;
}
const ThemeContext = React.createContext<ThemeContextType | null>(null);

// Concept 6: Reducer Pattern
type State = { count: number };
type Action = { type: "increment" } | { type: "decrement" };
function reducer(state: State, action: Action): State {
    switch (action.type) {
        case "increment": return { count: state.count + 1 };
        case "decrement": return { count: state.count - 1 };
        default: return state;
    }
}