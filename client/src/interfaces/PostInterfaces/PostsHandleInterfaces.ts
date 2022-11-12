//Post handler interfaces for classes

interface ResponseFromDelete {
    status: number;
    text: string;
}

export interface modifyUserPostsInterface {
    CheckIfPostDesIsToLong(): void;
    OpenModalEditHandler(): void;
    removePostHandler(): Promise<void | ResponseFromDelete>;
}

export interface HandlePostFormInterface {
    ResetValuesOfForm(): void;
    LogOut(): void;
    onFileSelected(e: any): void;
    GenerateDateString(): string;
    HandleSubmitPostForm(e: any): Promise<void>;
}