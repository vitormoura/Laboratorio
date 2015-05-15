using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;
using mopacman.Animations;
using mopacman.Controllers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    class Player : Sprite, IControllable
    {
        public event EventHandler ReadyToMove;

        public EnumDirections FacingDirection { get; set; }

        public MazeSection CurrentLocation
        {
            get { return this.currentLocation; }
            set
            {
                this.currentLocation = value;
                this.NextLocation = value;

                if (this.currentLocation != null)
                {
                    this.SetPosition(new Point(((int)(this.currentLocation.ID.X * this.Bounds.Width)), (int)(this.currentLocation.ID.Y * this.Bounds.Height)));
                    //this.OnReadyToMove();
                }
            }
        }

        public MazeSection PreviousLocation
        {
            get;
            set;
        }

        private MazeSection NextLocation
        {
            get;
            set;
        }

        protected DirectionalAnimation<Player> Animation
        {
            get
            {
                return this.animation;
            }
        }

        public Player(MyGame g, String assetName, Rectangle bounds)
            : base(g, assetName, bounds)
        {
            this.animation = new DirectionalAnimation<Player>(this, 0.25f);
            this.animation.Finished += animation_Finished;
        }

        public virtual void GoTo(EnumDirections d)
        {
            MazeSection next = null;

            if (this.CurrentLocation != null)
            {
                next = this.CurrentLocation.Get(d);

                if (next != null && next.Allowed)
                {
                    this.NextLocation = next;

                    if (d == EnumDirections.West || d == EnumDirections.East)
                        this.FacingDirection = d;

                    this.animation.Start(d, Constants.DEFAULT_BLOCK_WIDTH);

                    /*
                    this.PreviousLocation = this.CurrentLocation;
                    this.CurrentLocation = next;
                    */
                }
            }
        }

        public override void Update(GameTime gameTime)
        {
            this.animation.Update(gameTime);
            base.Update(gameTime);
        }

        public override void Draw(GameTime gameTime)
        {
            MyGame.SpriteBatch.Draw(this.Texture,
                destinationRectangle: MyGame.Camera.TranslateToPixelsRect(this.Bounds),
                effects: this.FacingDirection == EnumDirections.West ? SpriteEffects.FlipHorizontally : SpriteEffects.None);
        }

        private void animation_Finished(object sender, EventArgs e)
        {
            this.PreviousLocation = this.CurrentLocation;
            this.CurrentLocation = this.NextLocation;

            this.OnReadyToMove();
        }

        protected void OnReadyToMove()
        {
            if (this.ReadyToMove != null)
                this.ReadyToMove.Invoke(this, null);
        }

        private MazeSection currentLocation;
        private DirectionalAnimation<Player> animation;
    }
}
