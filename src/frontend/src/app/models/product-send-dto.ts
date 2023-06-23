export interface ProductSendDto {
  id: string;
  name: string;
  description: string | null | undefined;
  price: number;
  categoryId: string;
}
