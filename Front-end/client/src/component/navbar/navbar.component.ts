import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'navbar-component',
  templateUrl: './navbar.html',

  // satu paket yang dibawah
  standalone: true,
  imports: [CommonModule],
})
export class NavbarComponent {}
