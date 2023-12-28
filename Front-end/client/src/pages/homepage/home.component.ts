import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '../../component/navbar/navbar.component';
import { FooterComponent } from '../../component/footer/footer.component';
import { ComicFormat } from './home';

@Component({
  selector: 'home-page',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
  // satu paket yang dibawah
  standalone: true,
  imports: [CommonModule, NavbarComponent, FooterComponent],
})
export class HomePageComponent {
  arrComic: ComicFormat[] = [
    {
      name: 'jujutsu kaisen',
      imgUrl: 'https://wallpapercave.com/wp/wp10508782.jpg',
      price: 20000,
      deskripsi: 'balbalalblabllablalb',
      genre: 'action',
    },
    {
      name: 'jujutsu kaisen',
      imgUrl: 'https://wallpapercave.com/wp/wp10508782.jpg',
      price: 20000,
      deskripsi: 'balbalalblabllablalb',
      genre: 'action',
    },
    {
      name: 'jujutsu kaisen',
      imgUrl: 'https://wallpapercave.com/wp/wp10508782.jpg',
      price: 20000,
      deskripsi: 'balbalalblabllablalb',
      genre: 'action',
    },
    {
      name: 'jujutsu kaisen',
      imgUrl: 'https://wallpapercave.com/wp/wp10508782.jpg',
      price: 20000,
      deskripsi: 'balbalalblabllablalb',
      genre: 'action',
    },
    {
      name: 'jujutsu kaisen',
      imgUrl: 'https://wallpapercave.com/wp/wp10508782.jpg',
      price: 20000,
      deskripsi: 'balbalalblabllablalb',
      genre: 'action',
    },
    {
      name: 'jujutsu kaisen',
      imgUrl: 'https://wallpapercave.com/wp/wp10508782.jpg',
      price: 20000,
      deskripsi: 'balbalalblabllablalb',
      genre: 'action',
    },
    {
      name: 'jujutsu kaisen',
      imgUrl: 'https://wallpapercave.com/wp/wp10508782.jpg',
      price: 20000,
      deskripsi: 'balbalalblabllablalb',
      genre: 'action',
    },
    {
      name: 'jujutsu kaisen',
      imgUrl: 'https://wallpapercave.com/wp/wp10508782.jpg',
      price: 20000,
      deskripsi: 'balbalalblabllablalb',
      genre: 'action',
    },
  ];
}
