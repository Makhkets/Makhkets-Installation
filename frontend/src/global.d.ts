declare global {
    interface Window {
        updateProgress: (value: number) => void;
    }
}