import { Category } from "src/app/models/category";

export interface ProductState {
  category: Category | null;
  error: string | null;
}
