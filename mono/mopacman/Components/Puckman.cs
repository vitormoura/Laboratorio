﻿using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Audio;
using Microsoft.Xna.Framework.Graphics;
using Microsoft.Xna.Framework.Media;
using mopacman.Controllers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    class Puckman : Player
    {
        private Int32 frameIndex;
        private Double elapsedTime;
        private SoundEffect chompSound;
                        
        public Puckman(MyGame g)
            : base(g, "puckman.png", new Rectangle(0, 0, Constants.DEFAULT_BLOCK_WIDTH, Constants.DEFAULT_BLOCK_WIDTH))
        {
            this.frameIndex = 0;
            this.elapsedTime = 0.0;
        }

        protected override void LoadContent()
        {
            //this.chompSound = this.Game.Content.Load<SoundEffect>("pacman_chomp.wav");

            base.LoadContent();
        }

        public override void Draw(GameTime gameTime)
        {
            SpriteBatch sb = this.Game.Services.GetService<SpriteBatch>();
            SpriteEffects effects = SpriteEffects.None;
            float rotation = 0.0f;

            this.elapsedTime += gameTime.ElapsedGameTime.TotalSeconds;

            if( FacingDirection == EnumDirections.West )
            {
                effects = SpriteEffects.FlipHorizontally;
            }
            

            if (elapsedTime >= 0.2)
            {
                this.frameIndex = (this.frameIndex + 1) > 2 ? 0 : ++this.frameIndex;
                this.elapsedTime = 0.0;
            }

            Rectangle rectangleToDraw = new Rectangle(Constants.DEFAULT_BLOCK_WIDTH * frameIndex, 0, Constants.DEFAULT_BLOCK_WIDTH, Constants.DEFAULT_BLOCK_WIDTH);

            sb.Draw( this.Texture, destinationRectangle: this.Bounds, sourceRectangle: rectangleToDraw, effects: effects, rotation:rotation );
        }

        public override void GoTo(EnumDirections d)
        {
            base.GoTo(d);

            if (this.CurrentLocation.HasCookie)
            {
                this.CurrentLocation.HasCookie = false;
                //this.chompSound.Play();
            }
        }
    }
}
