import { Category } from "src/app/models/category";

export interface CategoryState {
  category: Category | null;
  error: string | null;
}
