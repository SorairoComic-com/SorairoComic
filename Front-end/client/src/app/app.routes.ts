import { Routes } from '@angular/router';
import { HomePageComponent } from '../pages/homepage/home.component';
import { TestComponent } from '../pages/testPage/test.component';

export const routes: Routes = [
  { path: '', component: HomePageComponent },
  { path: 'test', component: TestComponent },
  // { path: 'home', component: HomePageComponent },
  // Tambahkan rute lain jika diperlukan
];
