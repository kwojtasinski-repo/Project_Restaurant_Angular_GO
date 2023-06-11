export interface ProductSendDto {
  id: number;
  name: string;
  description: string | null | undefined;
  price: number;
  categoryId: number;
}
