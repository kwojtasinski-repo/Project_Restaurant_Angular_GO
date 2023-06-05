import { AppComponent } from "./app.component";
import { CartsComponent } from "./components/carts/carts.component";
import { AddCategoryComponent } from "./components/category/add-category/add-category.component";
import { CategoriesComponent } from "./components/category/categories/categories.component";
import { EditCategoryComponent } from "./components/category/edit-category/edit-category.component";
import { FooterComponent } from "./components/footer/footer.component";
import { HeaderComponent } from "./components/header/header.component";
import { LoginComponent } from "./components/login/login.component";
import { MenuComponent } from "./components/menu/menu.component";
import { MyOrdersComponent } from "./components/orders/my-orders/my-orders.component";
import { OrderViewComponent } from "./components/orders/order-view/order-view.component";
import { AddProductsComponent } from "./components/product/add-products/add-products.component";
import { EditProductsComponent } from "./components/product/edit-products/edit-products.component";
import { ProductFormComponent } from "./components/product/product-form/product-form.component";
import { ViewProductsComponent } from "./components/product/view-products/view-products.component";
import { RegisterComponent } from "./components/register/register.component";
import { SearchBarComponent } from "./components/search-bar/search-bar.component";
import { SpinnerButtonComponent } from "./components/spinner-button/spinner-button.component";
import { CurrencyFormatterDirective } from "./directives/currency-formatter-directive";
import { LoginFormDirective } from "./directives/login-form-directive";
import { MoneyPipe } from "./pipes/money-pipe";
import { ErrorDialogComponent } from "./services/error-dialog/error-dialog.component";

export const appDeclarations = [
    AppComponent,
    CategoriesComponent,
    AddCategoryComponent,
    AddProductsComponent,
    ViewProductsComponent,
    LoginComponent,
    MenuComponent,
    LoginFormDirective,
    HeaderComponent,
    CurrencyFormatterDirective,
    EditProductsComponent,
    ProductFormComponent,
    FooterComponent,
    EditCategoryComponent,
    SearchBarComponent,
    CartsComponent,
    OrderViewComponent,
    MoneyPipe,
    MyOrdersComponent,
    SpinnerButtonComponent,
    ErrorDialogComponent,
    RegisterComponent,
  ]
