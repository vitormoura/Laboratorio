using Microsoft.Xna.Framework;
using mopacman.Controllers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    class Player : Sprite, IControllable
    {
        public MazeSection CurrentLocation
        {
            get { return this.currentLocation; }
            set
            {
                this.currentLocation = value;

                if (this.currentLocation != null)
                {
                    this.SetPosition(new Point(((int)this.currentLocation.ID.X * this.Bounds.Width), ((int)this.currentLocation.ID.Y * this.Bounds.Height)));
                }
            }
        }

        public MazeSection PreviousLocation
        {
            get;
            set;
        }

        public Player(MyGame g, String assetName, Rectangle bounds)
            : base(g, assetName, bounds)
        {
        }

        public void GoTo(EnumDirections d)
        {
            MazeSection next = null;

            if (this.CurrentLocation != null)
            {
                next = this.CurrentLocation.Get(d);

                if (next != null && next.Allowed)
                {
                    this.PreviousLocation = this.CurrentLocation;
                    this.CurrentLocation = next;


                }
            }
        }
        

        private MazeSection currentLocation;
    }
}
